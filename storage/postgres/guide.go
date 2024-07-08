package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	pb "github.com/Azizbek-Qodirov/hr_platform_evaluation_service/genprotos"
	"github.com/google/uuid"
)

type GuideStorage struct {
	db *sql.DB
}

func NewGuideStorage(db *sql.DB) *GuideStorage {
	return &GuideStorage{db: db}
}

func (g *GuideStorage) CreateGuide(req *pb.CreateGuideRequest) (*pb.Void, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO guides (id, title, content, author)
		VALUES ($1, $2, $3, $4)
	`
	_, err := g.db.Exec(query, id, req.Title, req.Content, req.Author)
	if err != nil {
		log.Printf("Failed to create guide: %v", err)
		return &pb.Void{Success: false, Message: err.Error()}, err
	}

	return &pb.Void{Success: true, Message: "Guide created successfully!"}, nil
}

func (g *GuideStorage) GetGuide(req *pb.GetGuideRequest) (*pb.GuideResponse, error) {
	query := `
		SELECT id, title, content, author
		FROM guides
		WHERE id = $1 and deleted_at=0
	`
	row := g.db.QueryRow(query, req.Id)

	var guide pb.Guide
	err := row.Scan(&guide.Id, &guide.Title, &guide.Content, &guide.Author)
	if err != nil {
		log.Printf("Failed to get guide: %v", err)
		return nil, err
	}

	return &pb.GuideResponse{
		Success: true,
		Message: "Guide retrieved successfully!",
		Guide:   &guide,
	}, nil
}

func (g *GuideStorage) UpdateGuide(req *pb.UpdateGuideRequest) (*pb.Void, error) {
	query := `
		UPDATE guides
		SET title = $1, content = $2
		WHERE id = $4
	`
	_, err := g.db.Exec(query, req.Title, req.Content, req.Id)
	if err != nil {
		log.Printf("Failed to update guide: %v", err)
		return &pb.Void{Success: false, Message: err.Error()}, err
	}

	return &pb.Void{Success: true, Message: "Guide updated successfully!"}, nil
}

func (g *GuideStorage) DeleteGuide(req *pb.DeleteGuideRequest) (*pb.Void, error) {
	query := `
		update guides
		set deleted_at=$2
		WHERE id = $1
	`
	_, err := g.db.Exec(query, req.Id, time.Now().Unix())
	if err != nil {
		log.Printf("Failed to delete guide: %v", err)
		return &pb.Void{Success: false, Message: err.Error()}, err
	}

	return &pb.Void{Success: true, Message: "Guide deleted successfully!"}, nil
}

func (g *GuideStorage) ListAllGuides(req *pb.ListAllGuidesRequest) (*pb.ListAllGuidesResponse, error) {
	count := 1
	var arr []interface{}
	query := `
		SELECT id, title, content, author
		FROM guides where deleted_at=0 
	`
	if len(req.Title) > 0 {
		query += fmt.Sprintf(" AND title=$%d", count)
		count++
		arr = append(arr, req.Title)
	}
	if len(req.Content) > 0 {
		query += fmt.Sprintf(" AND content=$%d", count)
		count++
		arr = append(arr, req.Content)
	}
	rows, err := g.db.Query(query, arr...)
	if err != nil {
		log.Printf("Failed to list all guides: %v", err)
		return nil, err
	}
	defer rows.Close()

	var guides []*pb.Guide
	for rows.Next() {
		var guide pb.Guide
		if err := rows.Scan(&guide.Id, &guide.Title, &guide.Content, &guide.Author); err != nil {
			return nil, err
		}
		guides = append(guides, &guide)
	}

	return &pb.ListAllGuidesResponse{
		Success: true,
		Message: "Guides retrieved successfully!",
		Guides:  guides,
	}, nil
}

func (g *GuideStorage) SearchGuides(req *pb.SearchGuidesRequest) (*pb.ListAllGuidesResponse, error) {
	query := `
		SELECT id, title, content, author
		FROM guides
		WHERE title ILIKE '%' || $1 || '%' OR content ILIKE '%' || $2 || '%'
	`
	rows, err := g.db.Query(query, req.Title, req.Content)
	if err != nil {
		log.Printf("Failed to search guides: %v", err)
		return nil, err
	}
	defer rows.Close()

	var guides []*pb.Guide
	for rows.Next() {
		var guide pb.Guide
		if err := rows.Scan(&guide.Id, &guide.Title, &guide.Content, &guide.Author); err != nil {
			return nil, err
		}
		guides = append(guides, &guide)
	}

	return &pb.ListAllGuidesResponse{
		Success: true,
		Message: "Guides found successfully!",
		Guides:  guides,
	}, nil
}
