package show

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
	"time"

	"github.com/skiarn/browse-shows/templator"
)

//RemoteAPIURL implements templator.BuildAPIURL
func RemoteAPIURL(url *string) templator.BuildAPIURL {
	return func(params []string) string {
		return *url + "/site/programs/" + params[2]
	}
}

//Render implements templator.RenderTemplate
func Render() templator.RenderTemplate {
	return func(tmpl *template.Template, w http.ResponseWriter, tmplName string, body io.Reader) error {
		var info Show
		if err := json.NewDecoder(body).Decode(&info); err != nil {
			return err
		}
		return tmpl.ExecuteTemplate(w, tmplName+".html", info)
	}
}

//Show is a entity showing information about a show.
type Show struct {
	Nid                 string `json:"nid"`
	Name                string `json:"name"`
	Channel             string `json:"channel"`
	ChannelID           int    `json:"channel_id"`
	Description         string `json:"description"`
	ProgramImage        string `json:"program_image"`
	TopImage            string `json:"top_image"`
	Cycle               bool   `json:"cycle"`
	DarkTitle           bool   `json:"dark_title"`
	FacebookPageURL     string `json:"facebook_page_url"`
	HideTitle           bool   `json:"hide_title"`
	InstagramPageURL    string `json:"instagram_page_url"`
	ParticipantTabText  string `json:"participant_tab_text"`
	PrimeProgram        bool   `json:"prime_program"`
	ProgramImageVersion int    `json:"program_image_version"`
	ProgramTag          string `json:"program_tag"`
	ShowVideoBlock      bool   `json:"show_video_block"`
	TopImageVersion     int    `json:"top_image_version"`
	TwitterPageURL      string `json:"twitter_page_url"`
	YoutubePageURL      string `json:"youtube_page_url"`
	TopBlock            struct {
		Teasers []struct {
			ForReuse      bool        `json:"for_reuse"`
			Heading       string      `json:"heading"`
			ImageVersion  int         `json:"image_version"`
			Lead          string      `json:"lead"`
			LinkText      interface{} `json:"link_text"`
			LinkURL       string      `json:"link_url"`
			OffTime       time.Time   `json:"off_time"`
			OnTime        time.Time   `json:"on_time"`
			Position      int         `json:"position"`
			State         string      `json:"state"`
			UseOffTime    bool        `json:"use_off_time"`
			UseOnTime     bool        `json:"use_on_time"`
			Image         string      `json:"image"`
			MediaResource struct {
				AspectRatio interface{} `json:"aspect_ratio"`
				Byline      string      `json:"byline"`
				Description string      `json:"description"`
				Height      interface{} `json:"height"`
				Media       struct {
					URL interface{} `json:"url"`
				} `json:"media"`
				MediaMd5     interface{} `json:"media_md5"`
				Name         string      `json:"name"`
				UploadedByID interface{} `json:"uploaded_by_id"`
				Width        interface{} `json:"width"`
			} `json:"media_resource"`
		} `json:"teasers"`
	} `json:"top_block"`
	SocialBlockTeasers []struct {
		ArticleID     string `json:"article_id"`
		Heading       string `json:"heading"`
		ImageVersion  int    `json:"image_version"`
		Kicker        string `json:"kicker"`
		Lead          string `json:"lead"`
		LinkText      string `json:"link_text"`
		LinkURL       string `json:"link_url"`
		Position      int    `json:"position"`
		Type          string `json:"type"`
		Image         string `json:"image"`
		MediaResource struct {
			AspectRatio float64 `json:"aspect_ratio"`
			Byline      string  `json:"byline"`
			Description string  `json:"description"`
			Height      int     `json:"height"`
			Media       struct {
				URL string `json:"url"`
			} `json:"media"`
			MediaMd5     interface{} `json:"media_md5"`
			Name         string      `json:"name"`
			UploadedByID string      `json:"uploaded_by_id"`
			Width        int         `json:"width"`
		} `json:"media_resource"`
	} `json:"social_block_teasers"`
	LiveEvents  []interface{} `json:"live_events"`
	WeeksVideos struct {
		TopVideosHeading string `json:"top_videos_heading"`
		TopVideos        []int  `json:"top_videos"`
	} `json:"weeks_videos"`
	ParticipantGroups []struct {
		Name         string `json:"name"`
		Participants []struct {
			Description  string `json:"description"`
			ImageVersion int    `json:"image_version"`
			Name         string `json:"name"`
			PersonTag    string `json:"person_tag"`
			Image        struct {
				URL string `json:"url"`
			} `json:"image"`
		} `json:"participants"`
	} `json:"participant_groups"`
	CopyrightBlock struct {
		Text  string `json:"text"`
		Image string `json:"image"`
	} `json:"copyright_block"`
	VideoPoll struct {
	} `json:"video_poll"`
	SponsorContext struct {
		SponsorBlocks struct {
			Sponsor struct {
				Name            string `json:"name"`
				ForMetadata     bool   `json:"for_metadata"`
				SponsorElements []struct {
					LinkText    string `json:"link_text"`
					LinkURL     string `json:"link_url"`
					SponsorName string `json:"sponsor_name"`
					Image       struct {
						URL string `json:"url"`
					} `json:"image"`
				} `json:"sponsor_elements"`
			} `json:"sponsor"`
			Partner struct {
				Name            string        `json:"name"`
				SponsorElements []interface{} `json:"sponsor_elements"`
			} `json:"partner"`
			LocalPartner struct {
				Name            string        `json:"name"`
				SponsorElements []interface{} `json:"sponsor_elements"`
			} `json:"local_partner"`
			Supplier struct {
				Name            string        `json:"name"`
				SponsorElements []interface{} `json:"sponsor_elements"`
			} `json:"supplier"`
		} `json:"sponsor_blocks"`
	} `json:"sponsor_context"`
	NextBroadcastText string `json:"next_broadcast_text"`
	TwitterBlock      struct {
		SearchTerms string `json:"search_terms"`
	} `json:"twitter_block"`
	ProgramType          string `json:"program_type"`
	VideoQuickSelections []struct {
		Name string `json:"name"`
		Tag  string `json:"tag"`
	} `json:"video_quick_selections"`
}
