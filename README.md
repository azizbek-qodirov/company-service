# Emergency Response Service\


count := 1
	var arr []interface{}
	query := `
		SELECT id, evaluator_id, employee_id, rating
		FROM evaluations
		WHERE deleted_at=0 
	`
	if len(req.EmployeeId) > 0 {
		query += fmt.Sprintf(" and user_id=$%d", count)
		count++
		arr = append(arr, req.EmployeeId)
	}
	if len(req.EvaluatorId) > 0 {
		query += fmt.Sprintf(" and uvaluter_id=$%d", count)
		count++
		arr = append(arr, req.EvaluatorId)
	}

	rows, err := e.db.Query(query, arr...)