package postgres

import (
	"database/sql"
	"fmt"
	"time"

	pb "github.com/Azizbek-Qodirov/hr_platform_evaluation_service/genprotos"
)

type NotificationStorage struct {
	db *sql.DB
}

func NewNotificationStorage(db *sql.DB) *NotificationStorage {
	return &NotificationStorage{db: db}
}

// Create a new notification
func (n *NotificationStorage) Create(notification *pb.CreateNotification) (*pb.Void, error) {
	query := `
        INSERT INTO notifications (user_id, type, message, created_at)
        VALUES ($1, $2, $3, $4, $5)
    `
	_, err := n.db.Exec(query, notification.UserId, notification.Type, notification.Message, time.Now())
	if err != nil {
		return &pb.Void{Success: false, Message: err.Error()}, err
	}
	return &pb.Void{Success: true, Message: "Notification created successfully!"}, nil
}

// Get notifications by user ID
func (n *NotificationStorage) GetByUserId(request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	count := 1
	var arr []interface{}
	query := `
        SELECT user_id, type, message
        FROM notifications 
        deleted_at=0
    `
	if len(request.UserId) > 0 {
		query += fmt.Sprintf(" and user_id=$%d", count)
		count++
		arr = append(arr, request.UserId)
	}

	if len(request.NotifiationId) > 0 {
		query += fmt.Sprintf(" and id=$%d", count)
		count++
		arr = append(arr, request.NotifiationId)
	}
	rows, err := n.db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []*pb.CreateNotification
	for rows.Next() {
		var notification pb.CreateNotification
		err := rows.Scan(&notification.UserId, &notification.Type, &notification.Message)
		if err != nil {
			return nil, err
		}
		notifications=append(notifications, &notification)
	}
	return &pb.GetAllResponse{Notifications: notifications, Success: true}, nil
}

// Mark all notifications as read for a user
func (n *NotificationStorage) ReadeAll(request *pb.ReadeAllRequest) (*pb.Void, error) {
	query := `
        UPDATE notifications
        SET is_read = true
        WHERE user_id = $1
    `
	_, err := n.db.Exec(query, request.UserId)
	if err != nil {
		return &pb.Void{Success: false, Message: err.Error()}, err
	}
	return &pb.Void{Success: true, Message: "All notifications marked as read!"}, nil
}

// Send a notification to all users in a company
func (n *NotificationStorage) SendAll(request *pb.SendByCompanyidToUsers) (*pb.Void, error) {
	query := `
        select u.id from users  u
		join positions p on 
		u.position_id= p.id
		join department d on
		d.id=p.department_id
		join company c on
		c.id=d.company_id
		where company_id=$1
    `
	row, err := n.db.Query(query, request.CompanyId)
	if err != nil {
		return &pb.Void{Success: false, Message: err.Error()}, err
	}
	for row.Next() {
		var userid string
		err = row.Scan(&userid)
		if err != nil {
			return &pb.Void{Success: false, Message: err.Error()}, err
		}

		query = `
        INSERT INTO notifications (user_id, type, message, is_read, created_at)
        SELECT id, $1, $2, false, $3
        FROM users
        WHERE company_id = $4
    `
		_, err = n.db.Exec(query, userid, request.Type, request.Message, time.Now(), request.CompanyId)
		if err != nil {
			return &pb.Void{Success: false, Message: err.Error()}, err
		}

	}
	return &pb.Void{Success: true, Message: "Notification sent to all users!"}, nil
}




func (n *NotificationStorage) SendAllUsers(request *pb.SendByCompanyidToUsers) (*pb.Void, error) {
	query := `
        select id from users
    `
	row, err := n.db.Query(query)
	if err != nil {
		return &pb.Void{Success: false, Message: err.Error()}, err
	}
	for row.Next() {
		var userid string
		err = row.Scan(&userid)
		if err != nil {
			return &pb.Void{Success: false, Message: err.Error()}, err
		}

		query = `
        INSERT INTO notifications (user_id, type, message, is_read, created_at)
        SELECT id, $1, $2, false, $3
        FROM users
        WHERE company_id = $4
    `
		_, err = n.db.Exec(query, userid, request.Type, request.Message, time.Now(), request.CompanyId)
		if err != nil {
			return &pb.Void{Success: false, Message: err.Error()}, err
		}

	}
	return &pb.Void{Success: true, Message: "Notification sent to all users!"}, nil
}
