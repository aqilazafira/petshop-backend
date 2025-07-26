package config

import (
	"os"

	storage "github.com/supabase-community/storage-go"
)

var SupabaseClient *storage.Client

func InitSupabase() {
	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")

	SupabaseClient = storage.NewClient(supabaseURL, supabaseKey, map[string]string{
		"Authorization": "Bearer " + supabaseKey,
		"apikey":        supabaseKey,
	})
}
