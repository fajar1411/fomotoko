package redist

import (
	"time"

	"github.com/go-redis/redis/v8"
)

const Expiration = 1 * time.Hour

func RedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Sesuaikan dengan alamat Redis Anda
		Password: "",               // Kosongkan jika tidak ada password
		DB:       0,                // Gunakan database ke-0
	})
	return client
}

// CloseRedisConnection adalah fungsi untuk menutup koneksi ke Redis
func CloseRedisConnection(client *redis.Client) {
	client.Close()
}
