package jobs

import (
	"context"
	"log"
	"time"

	"gorm.io/gorm"
)

// ============================
// CLEANUP JOB ENTRY
// ============================
func StartCleanupJob(ctx context.Context, db *gorm.DB) {
	ticker := time.NewTicker(1 * time.Hour) // jalan tiap 1 jam

	go func() {
		log.Println("🧹 Cleanup Job started")

		for {
			select {
			case <-ticker.C:
				runCleanup(db)

			case <-ctx.Done():
				log.Println("🛑 Cleanup Job stopped")
				ticker.Stop()
				return
			}
		}
	}()
}

// ============================
// CLEANUP TASKS
// ============================
func runCleanup(db *gorm.DB) {
	log.Println("🧹 Running cleanup task...")

	cleanupRefreshTokens(db)
	cleanupOldActivityLogs(db)
	cleanupOldAuditLogs(db)

	log.Println("✅ Cleanup finished")
}

// ============================
// REFRESH TOKEN CLEANUP
// ============================
func cleanupRefreshTokens(db *gorm.DB) {
	result := db.Exec(`
		DELETE FROM refresh_tokens 
		WHERE expires_at < NOW() OR revoked = true
	`)

	if result.Error != nil {
		log.Println("❌ Failed cleanup refresh_tokens:", result.Error)
		return
	}

	log.Printf("🧹 Refresh tokens cleaned: %d rows\n", result.RowsAffected)
}

// ============================
// ACTIVITY LOG CLEANUP
// ============================
func cleanupOldActivityLogs(db *gorm.DB) {
	result := db.Exec(`
		DELETE FROM activity_logs
		WHERE created_at < NOW() - INTERVAL '30 days'
	`)

	if result.Error != nil {
		log.Println("❌ Failed cleanup activity_logs:", result.Error)
		return
	}

	log.Printf("🧹 Activity logs cleaned: %d rows\n", result.RowsAffected)
}

// ============================
// AUDIT LOG CLEANUP
// ============================
func cleanupOldAuditLogs(db *gorm.DB) {
	result := db.Exec(`
		DELETE FROM audit_logs
		WHERE created_at < NOW() - INTERVAL '90 days'
	`)

	if result.Error != nil {
		log.Println("❌ Failed cleanup audit_logs:", result.Error)
		return
	}

	log.Printf("🧹 Audit logs cleaned: %d rows\n", result.RowsAffected)
}
