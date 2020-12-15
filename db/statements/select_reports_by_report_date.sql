SELECT * FROM maude_report mr
INNER JOIN device d ON d.id = mr.device_id_fk
WHERE report_date >= $1 AND report_date < $2
