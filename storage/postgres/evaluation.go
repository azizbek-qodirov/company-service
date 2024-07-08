package postgres

import (
	"database/sql"
	"fmt"
	"log"

	pb "github.com/Azizbek-Qodirov/hr_platform_evaluation_service/genprotos"
	"github.com/google/uuid"
)

type EvaluationStorage struct {
	db *sql.DB
}

func NewEvaluationStorage(db *sql.DB) *EvaluationStorage {
	return &EvaluationStorage{db: db}
}

func (e *EvaluationStorage) Create(req *pb.EvaluationCreate) (*pb.Void, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO evaluations (id, evaluator_id, employee_id, rating, comment)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := e.db.Exec(query, id, req.EvaluatorId, req.EmployeeId, req.Rating, req.Comment)
	if err != nil {
		log.Printf("Failed to create evaluation: %v", err)
		return &pb.Void{Success: false, Message: err.Error()}, err
	}

	return &pb.Void{Success: true, Message: "Evaluation created successfully!"}, nil
}

func (e *EvaluationStorage) Get(req *pb.Byid) (*pb.EvaluationUpdate, error) {
	query := `
		SELECT id, evaluator_id, employee_id, rating, comment
		FROM evaluations
		WHERE id = $1
	`
	row := e.db.QueryRow(query, req.Id)

	var evaluation pb.EvaluationUpdate
	err := row.Scan(&evaluation.Id, &evaluation.EvaluatorId, &evaluation.EmployeeId, &evaluation.Rating, &evaluation.Comment)
	if err != nil {
		log.Printf("Failed to get evaluation: %v", err)
		return nil, err
	}

	return &evaluation, nil
}

func (e *EvaluationStorage) Update(req *pb.EvaluationUpdate) (*pb.Void, error) {
	query := `
		UPDATE evaluations
		SET evaluator_id = $1, employee_id = $2, rating = $3, comment = $4
		WHERE id = $5
	`
	_, err := e.db.Exec(query, req.EvaluatorId, req.EmployeeId, req.Rating, req.Comment, req.Id)
	if err != nil {
		log.Printf("Failed to update evaluation: %v", err)
		return &pb.Void{Success: false, Message: err.Error()}, err
	}

	return &pb.Void{Success: true, Message: "Evaluation updated successfully!"}, nil
}

func (e *EvaluationStorage) Delete(req *pb.Byid) (*pb.Void, error) {
	query := `
		DELETE FROM evaluations
		WHERE id = $1
	`
	_, err := e.db.Exec(query, req.Id)
	if err != nil {
		log.Printf("Failed to delete evaluation: %v", err)
		return &pb.Void{Success: false, Message: err.Error()}, err
	}

	return &pb.Void{Success: true, Message: "Evaluation deleted successfully!"}, nil
}

func (e *EvaluationStorage) GetAll(req *pb.EvaluationGetAllReq) (*pb.EvaluationGetAllRes, error) {
	count := 1
	var arr []interface{}
	query := `
		SELECT id, evaluator_id, employee_id, rating, comment
		FROM evaluations
		WHERE deleted_at IS NULL
	`
	if len(req.EmployeeId) > 0 {
		query += fmt.Sprintf(" AND employee_id=$%d", count)
		count++
		arr = append(arr, req.EmployeeId)
	}
	if len(req.EvaluatorId) > 0 {
		query += fmt.Sprintf(" AND evaluator_id=$%d", count)
		count++
		arr = append(arr, req.EvaluatorId)
	}

	rows, err := e.db.Query(query, arr...)
	if err != nil {
		log.Printf("Failed to get all evaluations: %v", err)
		return nil, err
	}
	defer rows.Close()

	var evaluations []*pb.EvaluationUpdate
	for rows.Next() {
		var evaluation pb.EvaluationUpdate
		if err := rows.Scan(&evaluation.Id, &evaluation.EvaluatorId, &evaluation.EmployeeId, &evaluation.Rating, &evaluation.Comment); err != nil {
			return nil, err
		}
		evaluations = append(evaluations, &evaluation)
	}

	return &pb.EvaluationGetAllRes{
		Evaluations: evaluations,
		Success:     true,
	}, nil
}
