package main

type InstagramDto struct {
	Config struct {
		CsrfToken string      `json:"csrf_token"`
		Viewer    interface{} `json:"viewer"`
		ViewerID  interface{} `json:"viewerId"`
	} `json:"config"`
	CountryCode  string `json:"country_code"`
	LanguageCode string `json:"language_code"`
	Locale       string `json:"locale"`
	EntryData    struct {
		ProfilePage []struct {
			LoggingPageID         string `json:"logging_page_id"`
			ShowSuggestedProfiles bool   `json:"show_suggested_profiles"`
			Graphql               struct {
				User struct {
					Biography              string `json:"biography"`
					BlockedByViewer        bool   `json:"blocked_by_viewer"`
					CountryBlock           bool   `json:"country_block"`
					ExternalURL            string `json:"external_url"`
					ExternalURLLinkshimmed string `json:"external_url_linkshimmed"`
					EdgeFollowedBy         struct {
						Count int `json:"count"`
					} `json:"edge_followed_by"`
					FollowedByViewer bool `json:"followed_by_viewer"`
					EdgeFollow       struct {
						Count int `json:"count"`
					} `json:"edge_follow"`
					FollowsViewer        bool   `json:"follows_viewer"`
					FullName             string `json:"full_name"`
					HasChannel           bool   `json:"has_channel"`
					HasBlockedViewer     bool   `json:"has_blocked_viewer"`
					HighlightReelCount   int    `json:"highlight_reel_count"`
					HasRequestedViewer   bool   `json:"has_requested_viewer"`
					ID                   string `json:"id"`
					IsBusinessAccount    bool   `json:"is_business_account"`
					IsJoinedRecently     bool   `json:"is_joined_recently"`
					BusinessCategoryName string `json:"business_category_name"`
					IsPrivate            bool   `json:"is_private"`
					IsVerified           bool   `json:"is_verified"`
					EdgeMutualFollowedBy struct {
						Count int           `json:"count"`
						Edges []interface{} `json:"edges"`
					} `json:"edge_mutual_followed_by"`
					ProfilePicURL          string      `json:"profile_pic_url"`
					ProfilePicURLHd        string      `json:"profile_pic_url_hd"`
					RequestedByViewer      bool        `json:"requested_by_viewer"`
					Username               string      `json:"username"`
					ConnectedFbPage        interface{} `json:"connected_fb_page"`
					EdgeFelixVideoTimeline struct {
						Count    int `json:"count"`
						PageInfo struct {
							HasNextPage bool        `json:"has_next_page"`
							EndCursor   interface{} `json:"end_cursor"`
						} `json:"page_info"`
						Edges []interface{} `json:"edges"`
					} `json:"edge_felix_video_timeline"`
					EdgeOwnerToTimelineMedia struct {
						Count    int `json:"count"`
						PageInfo struct {
							HasNextPage bool   `json:"has_next_page"`
							EndCursor   string `json:"end_cursor"`
						} `json:"page_info"`
						Edges []struct {
							Node struct {
								Typename           string `json:"__typename"`
								ID                 string `json:"id"`
								EdgeMediaToCaption struct {
									Edges []struct {
										Node struct {
											Text string `json:"text"`
										} `json:"node"`
									} `json:"edges"`
								} `json:"edge_media_to_caption"`
								Shortcode          string `json:"shortcode"`
								EdgeMediaToComment struct {
									Count int `json:"count"`
								} `json:"edge_media_to_comment"`
								CommentsDisabled bool `json:"comments_disabled"`
								TakenAtTimestamp int  `json:"taken_at_timestamp"`
								Dimensions       struct {
									Height int `json:"height"`
									Width  int `json:"width"`
								} `json:"dimensions"`
								DisplayURL  string `json:"display_url"`
								EdgeLikedBy struct {
									Count int `json:"count"`
								} `json:"edge_liked_by"`
								EdgeMediaPreviewLike struct {
									Count int `json:"count"`
								} `json:"edge_media_preview_like"`
								Location     interface{} `json:"location"`
								GatingInfo   interface{} `json:"gating_info"`
								MediaPreview string      `json:"media_preview"`
								Owner        struct {
									ID       string `json:"id"`
									Username string `json:"username"`
								} `json:"owner"`
								ThumbnailSrc       string `json:"thumbnail_src"`
								ThumbnailResources []struct {
									Src          string `json:"src"`
									ConfigWidth  int    `json:"config_width"`
									ConfigHeight int    `json:"config_height"`
								} `json:"thumbnail_resources"`
								IsVideo              bool   `json:"is_video"`
								AccessibilityCaption string `json:"accessibility_caption"`
							} `json:"node,omitempty"`
						} `json:"edges"`
					} `json:"edge_owner_to_timeline_media"`
					EdgeSavedMedia struct {
						Count    int `json:"count"`
						PageInfo struct {
							HasNextPage bool        `json:"has_next_page"`
							EndCursor   interface{} `json:"end_cursor"`
						} `json:"page_info"`
						Edges []interface{} `json:"edges"`
					} `json:"edge_saved_media"`
					EdgeMediaCollections struct {
						Count    int `json:"count"`
						PageInfo struct {
							HasNextPage bool        `json:"has_next_page"`
							EndCursor   interface{} `json:"end_cursor"`
						} `json:"page_info"`
						Edges []interface{} `json:"edges"`
					} `json:"edge_media_collections"`
				} `json:"user"`
			} `json:"graphql"`
			ShowFollowDialog bool `json:"show_follow_dialog"`
		} `json:"ProfilePage"`
	} `json:"entry_data"`
	Hostname        string  `json:"hostname"`
	DeploymentStage string  `json:"deployment_stage"`
	Platform        string  `json:"platform"`
	RhxGis          string  `json:"rhx_gis"`
	Nonce           string  `json:"nonce"`
	MidPct          float64 `json:"mid_pct"`
	ZeroData        struct {
	} `json:"zero_data"`
	CacheSchemaVersion int `json:"cache_schema_version"`
	ServerChecks       struct {
	} `json:"server_checks"`
	Knobs struct {
		AcctNtb int `json:"acct:ntb"`
		Cb      int `json:"cb"`
		Captcha int `json:"captcha"`
		Fr      int `json:"fr"`
	} `json:"knobs"`
	ToCache struct {
		Gatekeepers struct {
			Num3  bool `json:"3"`
			Num4  bool `json:"4"`
			Num5  bool `json:"5"`
			Num6  bool `json:"6"`
			Num7  bool `json:"7"`
			Num8  bool `json:"8"`
			Num9  bool `json:"9"`
			Num10 bool `json:"10"`
			Num11 bool `json:"11"`
			Num12 bool `json:"12"`
			Num13 bool `json:"13"`
			Num14 bool `json:"14"`
			Num15 bool `json:"15"`
			Num16 bool `json:"16"`
			Num18 bool `json:"18"`
			Num19 bool `json:"19"`
			Num20 bool `json:"20"`
			Num21 bool `json:"21"`
			Num23 bool `json:"23"`
			Num24 bool `json:"24"`
			Num25 bool `json:"25"`
			Num26 bool `json:"26"`
			Num27 bool `json:"27"`
		} `json:"gatekeepers"`
		Qe struct {
			Num0 struct {
				P struct {
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"0"`
			Num2 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"2"`
			Num4 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"4"`
			Num5 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"5"`
			Num6 struct {
				P struct {
					Num0 string `json:"0"`
					Num1 bool   `json:"1"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"6"`
			Num7 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"7"`
			Num8 struct {
				P struct {
					Num0 bool `json:"0"`
					Num1 bool `json:"1"`
					Num2 bool `json:"2"`
					Num3 bool `json:"3"`
					Num4 bool `json:"4"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"8"`
			Num9 struct {
				P struct {
					Num0 bool `json:"0"`
					Num1 bool `json:"1"`
					Num2 int  `json:"2"`
					Num3 int  `json:"3"`
					Num4 int  `json:"4"`
					Num5 bool `json:"5"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"9"`
			Num10 struct {
				P struct {
					Num0 bool `json:"0"`
					Num1 bool `json:"1"`
					Num2 bool `json:"2"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"10"`
			Num11 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"11"`
			Num12 struct {
				P struct {
					Num0 int `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"12"`
			Sdc struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"sdc"`
			AppUpsell struct {
				G string `json:"g"`
				P struct {
					HasIgliteNewContent string `json:"has_iglite_new_content"`
					HasIgliteLink       string `json:"has_iglite_link"`
				} `json:"p"`
			} `json:"app_upsell"`
			IglAppUpsell struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"igl_app_upsell"`
			FrxReporting struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"frx_reporting"`
			AccRecoveryLink struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"acc_recovery_link"`
			Notif struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"notif"`
			ShowCopyLink struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"show_copy_link"`
			PEdit struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"p_edit"`
			AccRecovery struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"acc_recovery"`
			Bundle struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"bundle"`
			Collections struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"collections"`
			CommentTa struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"comment_ta"`
			Su struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"su"`
			EbdsimLi struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"ebdsim_li"`
			EbdsimLo struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"ebdsim_lo"`
			EmptyFeed struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"empty_feed"`
			Appsell struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"appsell"`
			HeartTab struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"heart_tab"`
			FollowButton struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"follow_button"`
			LogCont struct {
				G string `json:"g"`
				P struct {
					HasContextualM string `json:"has_contextual_m"`
				} `json:"p"`
			} `json:"log_cont"`
			Onetaplogin struct {
				G string `json:"g"`
				P struct {
					DuringReg      string `json:"during_reg"`
					DefaultValue   string `json:"default_value"`
					StorageVersion string `json:"storage_version"`
				} `json:"p"`
			} `json:"onetaplogin"`
			Onetap struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"onetap"`
			ProfileTabs struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"profile_tabs"`
			MultiregIter struct {
				G string `json:"g"`
				P struct {
					HasNewPhoneForm    string `json:"has_new_phone_form"`
					HasPhoneCodeFilter string `json:"has_phone_code_filter"`
				} `json:"p"`
			} `json:"multireg_iter"`
			RegVp struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"reg_vp"`
			SuUniverse struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"su_universe"`
			TpPblshr struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"tp_pblshr"`
			Felix struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"felix"`
			FelixClearFbCookie struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"felix_clear_fb_cookie"`
			FelixCreationDurationLimits struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"felix_creation_duration_limits"`
			FelixCreationEnabled struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"felix_creation_enabled"`
			FelixCreationFbCrossposting struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"felix_creation_fb_crossposting"`
			FelixCreationFbCrosspostingV2 struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"felix_creation_fb_crossposting_v2"`
			FelixCreationValidation struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"felix_creation_validation"`
			ProfileEnhanceLi struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"profile_enhance_li"`
			ProfileEnhanceLo struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"profile_enhance_lo"`
			CommentEnhance struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"comment_enhance"`
			MwebTopicalExplore struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"mweb_topical_explore"`
			FollowAllFb struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"follow_all_fb"`
			A2HsHeuristicNonUc struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"a2hs_heuristic_non_uc"`
			WebHashtag struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"web_hashtag"`
			WebHashtagLoggedOut struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"web_hashtag_logged_out"`
			HeaderScroll struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"header_scroll"`
			LiteRating struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"lite_rating"`
			WebEmbedsShare struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"web_embeds_share"`
			WebEmbedsLoggedOut struct {
				G string `json:"g"`
				P struct {
					ShowCommentInput string `json:"show_comment_input"`
				} `json:"p"`
			} `json:"web_embeds_logged_out"`
			WebDatasaverMode struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"web_datasaver_mode"`
			LiteDatasaverMode struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"lite_datasaver_mode"`
			PostOptions struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"post_options"`
			IgtvCreationFeedPreview struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"igtv_creation_feed_preview"`
			IgtvPublicViewing struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"igtv_public_viewing"`
			Nux struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"nux"`
			Iglmsr struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"iglmsr"`
			Igwsvl struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"igwsvl"`
			Iglcp struct {
				G string `json:"g"`
				P struct {
					HasLoginPrefill string `json:"has_login_prefill"`
				} `json:"p"`
			} `json:"iglcp"`
			Iglscioi struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"iglscioi"`
			Ws2 struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"ws2"`
			Wpn struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"wpn"`
			LiteUserTagCreation struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"lite_user_tag_creation"`
			Iglcius struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"iglcius"`
			Wfvu struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"wfvu"`
			Iglsa struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"iglsa"`
			Igwllre struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"igwllre"`
			Wss2 struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"wss2"`
			StickerTray struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"sticker_tray"`
			WebSentry struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"web_sentry"`
		} `json:"qe"`
		ProbablyHasApp bool `json:"probably_has_app"`
		Cb             bool `json:"cb"`
	} `json:"to_cache"`
	RolloutHash   string `json:"rollout_hash"`
	BundleVariant string `json:"bundle_variant"`
	IsCanary      bool   `json:"is_canary"`
}
