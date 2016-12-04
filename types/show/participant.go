package show

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
	"time"

	"github.com/skiarn/browse-shows/templator"
)

//RemoteParticipantAPIURL implements templator.BuildAPIURL.
func RemoteParticipantAPIURL(url *string) templator.BuildAPIURL {
	return func(params []string) string {
		return *url + "/play/video_assets.json?tags=" + params[2]
	}
}

//RenderParticipant implements templator.RenderTemplate for Participant.
func RenderParticipant(link *string) templator.RenderTemplate {
	return func(tmpl *template.Template, w http.ResponseWriter, tmplName string, body io.Reader) error {
		var info Participant
		if err := json.NewDecoder(body).Decode(&info); err != nil {
			return err
		}
		data := struct {
			Participant *Participant
			Link        string
		}{&info, *link}
		return tmpl.ExecuteTemplate(w, tmplName+"-participant.html", data)
	}
}

//Participant is a entity showing information about a participant in a show.
type Participant struct {
	TotalHits int `json:"total_hits"`
	Results   []struct {
		ID          string   `json:"id"`
		Title       string   `json:"title"`
		Description string   `json:"description"`
		IsLive      bool     `json:"is_live"`
		Tags        []string `json:"tags"`
		Image       string   `json:"image"`
		Program     struct {
			Airtime                 string        `json:"airtime"`
			CarouselImage           string        `json:"carousel_image"`
			CreatedAt               time.Time     `json:"created_at"`
			Description             string        `json:"description"`
			DisallowEmbed           bool          `json:"disallow_embed"`
			ExternalLinkText        string        `json:"external_link_text"`
			ExternalLinkURL         string        `json:"external_link_url"`
			FacebookPageName        string        `json:"facebook_page_name"`
			HideAds                 bool          `json:"hide_ads"`
			HideChannelLogoInPlayer bool          `json:"hide_channel_logo_in_player"`
			IsPremium               bool          `json:"is_premium"`
			IsRecommended           bool          `json:"is_recommended"`
			IsSponsored             bool          `json:"is_sponsored"`
			Logo                    string        `json:"logo"`
			MetaDescription         string        `json:"meta_description"`
			MetaTitle               string        `json:"meta_title"`
			Name                    string        `json:"name"`
			Nid                     string        `json:"nid"`
			Position                int           `json:"position"`
			ProgramImage            string        `json:"program_image"`
			ProgramTag              string        `json:"program_tag"`
			ShowSimilarPrograms     bool          `json:"show_similar_programs"`
			SimilarPrograms         []string      `json:"similar_programs"`
			UpdatedAt               time.Time     `json:"updated_at"`
			Genres                  []interface{} `json:"genres"`
			Channel                 struct {
				About string `json:"about"`
				Links []struct {
					Link string `json:"link"`
					Name string `json:"name"`
				} `json:"links"`
				Logo         string `json:"logo"`
				LogoVersion  int    `json:"logo_version"`
				Name         string `json:"name"`
				Nid          string `json:"nid"`
				Position     int    `json:"position"`
				VideoAssetID string `json:"video_asset_id"`
			} `json:"channel"`
			Category struct {
				Name     string `json:"name"`
				Nid      string `json:"nid"`
				Position int    `json:"position"`
				HideAds  bool   `json:"hide_ads"`
			} `json:"category"`
			Tags         []string `json:"tags"`
			OldProgramID string   `json:"old_program_id"`
			Burt         struct {
				Category string   `json:"category"`
				Tags     []string `json:"tags"`
			} `json:"burt"`
			CinemascopeImagePath string `json:"cinemascope_image_path"`
		} `json:"program"`
		ProgramNid   string `json:"program_nid"`
		Availability struct {
			AvailabilityGroupFree    string   `json:"availability_group_free"`
			AvailabilityGroupPremium string   `json:"availability_group_premium"`
			GeoRestricted            bool     `json:"geo_restricted"`
			Live                     bool     `json:"live"`
			Human                    string   `json:"human"`
			Groups                   []string `json:"groups"`
		} `json:"availability"`
		BroadcastDateTime time.Time   `json:"broadcast_date_time"`
		Categories        []string    `json:"categories"`
		Keywords          []string    `json:"keywords"`
		HideAds           bool        `json:"hide_ads"`
		ProgramTypeID     int         `json:"program_type_id"`
		IsClip            bool        `json:"is_clip"`
		FullProgram       bool        `json:"full_program"`
		AllowEmbed        bool        `json:"allow_embed"`
		Platform          string      `json:"platform"`
		Duration          int         `json:"duration"`
		ShowLogo          bool        `json:"show_logo"`
		ExternalID        interface{} `json:"external_id"`
		Hd                bool        `json:"hd"`
		IsChannel         bool        `json:"is_channel"`
		DateTimeIndexed   time.Time   `json:"date_time_indexed"`
		Akamai            struct {
			Channel          string `json:"channel"`
			VideoPlatform    string `json:"videoPlatform"`
			ContentID        string `json:"contentId"`
			ContentType      string `json:"contentType"`
			IsLive           string `json:"isLive"`
			Title            string `json:"title"`
			AssetName        string `json:"assetName"`
			Category         string `json:"category"`
			SubCategory      string `json:"subCategory"`
			Categories       string `json:"categories"`
			DeliveryType     string `json:"deliveryType"`
			ContentLength    string `json:"contentLength"`
			ProductGroupsIds string `json:"productGroupsIds"`
		} `json:"akamai"`
		Videoplaza struct {
			Shares      string `json:"shares"`
			ContentForm string `json:"contentForm"`
			Flags       string `json:"flags"`
			ContentID   string `json:"contentId"`
			Tags        string `json:"tags"`
		} `json:"videoplaza"`
		Platforms             []string      `json:"platforms"`
		PublishedDateTime     time.Time     `json:"published_date_time"`
		ExpireDateTime        interface{}   `json:"expire_date_time"`
		PremiumExpireDateTime interface{}   `json:"premium_expire_date_time"`
		IsDrmProtected        bool          `json:"is_drm_protected"`
		ProductGroups         []int         `json:"product_groups"`
		GeoRegions            []string      `json:"geo_regions"`
		IsGeoRestricted       bool          `json:"is_geo_restricted"`
		MediaFormats          []string      `json:"media_formats"`
		ProductGroupNids      []interface{} `json:"product_group_nids"`
		Burt                  struct {
			Category string   `json:"category"`
			Tags     []string `json:"tags"`
		} `json:"burt"`
	} `json:"results"`
	AssetsTypesHits struct {
		Clip int `json:"clip"`
	} `json:"assets_types_hits"`
}
