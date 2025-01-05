package repository

import (
	"fmt"
	"net/http"

	gol0 "github.com/FrenkenFlores/golang_l0"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SslMode  string
}

type PostgresDB struct {
	database *sqlx.DB
}

func (r *PostgresDB) GetOrder(id string) (int, map[string]any) {
	var order gol0.Order
	row := r.database.QueryRow("SELECT * FROM orders WHERE id=$1", id)
	row.Scan(
		&order.Id,
		&order.OrderUid,
		&order.TrackNumber,
		&order.Entry,
		&order.Locale,
		&order.InternalSignature,
		&order.CustomerId,
		&order.DeliveryService,
		&order.SharedKey,
		&order.SmId,
		&order.DateCreated,
		&order.OofShard,
	)
	row = r.database.QueryRow("SELECT * FROM delivery WHERE id=$1", id)
	row.Scan(
		&order.Delivery.Id,
		&order.Delivery.OrderId,
		&order.Delivery.Name,
		&order.Delivery.Phone,
		&order.Delivery.Zip,
		&order.Delivery.City,
		&order.Delivery.Address,
		&order.Delivery.Region,
		&order.Delivery.Email,
	)
	row = r.database.QueryRow("SELECT * FROM payment WHERE id=$1", id)
	row.Scan(
		&order.Payment.Id,
		&order.Payment.OrderId,
		&order.Payment.Transaction,
		&order.Payment.RequestId,
		&order.Payment.Currency,
		&order.Payment.Provider,
		&order.Payment.Amount,
		&order.Payment.PaymentDt,
		&order.Payment.Bank,
		&order.Payment.DeliveryCost,
		&order.Payment.GoodsTotal,
		&order.Payment.CustomFee,
	)
	rows, err := r.database.Query("SELECT * FROM items WHERE id=$1", id)
	if err != nil {
		return http.StatusInternalServerError, map[string]any{"error": err.Error()}
	}
	defer rows.Close()
	var item gol0.Item
	for rows.Next() {
		err := rows.Scan(
			&item.Id,
			&item.OrderId,
			&item.ChrtId,
			&item.TrackNumber,
			&item.Price,
			&item.Rid,
			&item.Name,
			&item.Sale,
			&item.Size,
			&item.TotalPrice,
			&item.NmId,
			&item.Brand,
			&item.Status,
		)
		if err != nil {
			return http.StatusInternalServerError, map[string]any{"error": err.Error()}
		}
		order.Items = append(order.Items, item)
	}
	res := make(map[string]interface{})
	{
		res["order_uid"] = order.OrderUid
		res["track_number"] = order.TrackNumber
		res["entry"] = order.Entry
		res["delivery"] = map[string]string{
			"name":    order.Delivery.Name,
			"phone":   order.Delivery.Phone,
			"zip":     order.Delivery.Zip,
			"city":    order.Delivery.City,
			"address": order.Delivery.Address,
			"region":  order.Delivery.Region,
			"email":   order.Delivery.Email,
		}
		res["payment"] = map[string]interface{}{
			"transaction":   order.Payment.Transaction,
			"request_id":    order.Payment.RequestId,
			"currency":      order.Payment.Currency,
			"provider":      order.Payment.Provider,
			"amount":        order.Payment.Amount,
			"payment_dt":    order.Payment.PaymentDt,
			"bank":          order.Payment.Bank,
			"delivery_cost": order.Payment.DeliveryCost,
			"goods_total":   order.Payment.GoodsTotal,
			"custom_fee":    order.Payment.CustomFee,
		}
		res["items"] = []map[string]interface{}{}
		for _, item := range order.Items {
			res["items"] = append(res["items"].([]map[string]interface{}), map[string]interface{}{
				"track_number": item.TrackNumber,
				"price":        item.Price,
				"rid":          item.Rid,
				"name":         item.Name,
				"sale":         item.Sale,
				"size":         item.Size,
				"total_price":  item.TotalPrice,
				"nm_id":        item.NmId,
				"brand":        item.Brand,
				"status":       item.Status,
			})
		}
		res["locale"] = order.Locale
		res["internal_signature"] = order.InternalSignature
		res["customer_id"] = order.CustomerId
		res["delivery_service"] = order.DeliveryService
		res["shared_key"] = order.SharedKey
		res["sm_id"] = order.SmId
		res["date_created"] = order.DateCreated
		res["oof_shard"] = order.OofShard
	}
	return http.StatusOK, res
}

func NewPostgresDb(cfg Config) (*PostgresDB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SslMode,
	))
	if err != nil {

		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &PostgresDB{database: db}, err
}
