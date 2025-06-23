package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware adalah middleware untuk memeriksa token autentikasi
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Ambil token dari header Authorization
        token := c.GetHeader("Authorization")

        // Cek apakah token ada
        if token == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
            c.Abort() // Hentikan eksekusi lebih lanjut
            return
        }

        // Di sini Anda dapat menambahkan logika untuk memverifikasi token
        // Misalnya, memeriksa apakah token valid atau tidak
        // Untuk contoh ini, kita hanya akan memeriksa apakah token adalah "valid-token"
        if token != "valid-token" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
            c.Abort() // Hentikan eksekusi lebih lanjut
            return
        }

        // Jika token valid, lanjutkan ke handler berikutnya
        c.Next()
    }
}