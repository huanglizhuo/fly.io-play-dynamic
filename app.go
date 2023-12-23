package main

import (
	"embed"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
)

type App struct {
	Relation []string `json:"relation"`
	Target   struct {
		Namespace              string   `json:"namespace"`
		PackageName            string   `json:"package_name"`
		SHA256CertFingerprints []string `json:"sha256_cert_fingerprints"`
	} `json:"target"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Create an array of App instances with the specified JSON data
	apps := []App{
		{
			Relation: []string{"delegate_permission/common.handle_all_urls"},
			Target: struct {
				Namespace              string   `json:"namespace"`
				PackageName            string   `json:"package_name"`
				SHA256CertFingerprints []string `json:"sha256_cert_fingerprints"`
			}{
				Namespace:              "android_app",
				PackageName:            "com.nttdocomo.android.mymagazine.debug",
				SHA256CertFingerprints: []string{"96:DF:74:F7:CD:0D:B9:59:6B:08:4A:B2:B4:4D:68:7E:0E:17:0D:BE:15:86:28:74:0B:E5:2F:3D:9B:AF:23:7C"},
			},
		},
		{
			Relation: []string{"delegate_permission/common.handle_all_urls"},
			Target: struct {
				Namespace              string   `json:"namespace"`
				PackageName            string   `json:"package_name"`
				SHA256CertFingerprints []string `json:"sha256_cert_fingerprints"`
			}{
				Namespace:              "android_app",
				PackageName:            "jp.gocro.smartnews.android.debug",
				SHA256CertFingerprints: []string{"96:DF:74:F7:CD:0D:B9:59:6B:08:4A:B2:B4:4D:68:7E:0E:17:0D:BE:15:86:28:74:0B:E5:2F:3D:9B:AF:23:7C"},
			},
		},
	}

	jsonData, err := json.MarshalIndent(apps, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Set the Content-Type header and write the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func dynamicLinkHandler(w http.ResponseWriter, r *http.Request) {
	// Redirect to the specified URL for the "dynamicLink" route
	http.Redirect(w, r, "https://ean5.adj.st/openArticle?placement=article-preview&identifier=4626881932293047673&adj_t=ejmhr4_8srfzb", http.StatusTemporaryRedirect)
}

func sfdLinkHandler(w http.ResponseWriter, r *http.Request) {
	// Redirect to the specified URL for the "dynamicLink" route
	http.Redirect(w, r, "https://ean5.adj.st/openArticle?placement=article-preview&identifier=4626881932293047673&adj_t=ejmhr4_8srfzb", http.StatusTemporaryRedirect)
}

func snLinkHandler(w http.ResponseWriter, r *http.Request) {
	// Redirect to the specified URL for the "dynamicLink" route
	http.Redirect(w, r, "https://ean5.adj.st/openArticle?placement=article-preview&identifier=4626881932293047673&adj_t=ejmhr4_8srfzb", http.StatusTemporaryRedirect)
}

//go:embed templates/*
var resources embed.FS

var t = template.Must(template.ParseFS(resources, "templates/*"))

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"

	}

	http.HandleFunc("/.well-known/assetlinks.json", handler)
	http.HandleFunc("/dynamicLink", dynamicLinkHandler)
	http.HandleFunc("/sfdLink", sfdLinkHandler)
	http.HandleFunc("/snLink", snLinkHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"Region": os.Getenv("FLY_REGION"),
		}

		t.ExecuteTemplate(w, "index.html.tmpl", data)
	})

	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
