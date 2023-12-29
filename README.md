# Douyin SDK for Go

[![Build](https://img.shields.io/badge/github-passing-green?style=flat&logo=github)](https://github.com/guaidashu/douyin-go/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/guaidashu/douyin-go)](https://goreportcard.com/report/github.com/guaidashu/douyin-go)
[![Version](https://img.shields.io/github/release/guaidashu/douyin-go.svg?style=flat)](https://github.com/guaidashu/douyin-go/releases/latest)
[![Reference](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](http://godoc.org/github.com/guaidashu/douyin-go)
[![Licence](https://img.shields.io/github/license/guaidashu/douyin-go?style=flat)](https://github.com/guaidashu/douyin-go/blob/master/LICENSE)

抖音开放平台SDK

## 安装
```go
import douyinGo "github.com/guaidashu/douyin-go"
```
## 使用
**初始化**
```go
credentials := douyinGo.NewCredentials("CLIENT_KEY", "CLIENT_SECRET")
manager := douyinGo.NewManager(credentials, nil)
```

**生成授权链接,获取授权码** `/platform/oauth/connect/`
```go
oauthUrl := manager.OauthConnect(douyinGo.OauthParam{
    Scope: "user_info,mobile_alert,video.list,video.data,video.create,video.delete,data.external.user,data.external.item,aweme.share,fans.list,following.list,item.comment,star_top_score_display,fans.data,data.external.fans_source,data.external.fans_favourite,discovery.ent,video.search,video.search.comment,fans.check",
    RedirectUri: "REDIRECT_URI",
})
```

**获取AccessToken** `/oauth/access_token/`
```go
accessToken, err := manager.OauthAccessToken(douyinGo.OauthAccessTokenReq{
    Code: "CODE",
})
```

**刷新access_token** `/oauth/refresh_token/`
```go
manager.OauthRenewRefreshToken(douyinGo.OauthRenewRefreshTokenReq{
    RefreshToken: "REFRESH_TOKEN",
})
```

**刷新refresh_token** `/oauth/renew_refresh_token/`
```go
manager.OauthRenewRefreshToken(douyinGo.OauthRenewRefreshTokenReq{
    RefreshToken: "REFRESH_TOKEN",
})
```

**生成client_token** `/oauth/client_token/`
```go
clientToken, err := manager.OauthClientAccessToken()
```

**获取用户信息** `/oauth/userinfo/`
```go
userInfo, err := manager.OauthUserinfo(douyinGo.OauthUserinfoReq{
    OpenId:      "OPEN_ID",
    AccessToken: "ACCESS_TOKEN",
})

// 解析手机号
mobile, err := manager.DecryptMobile("ENCRYPT_MOBILE")
```

**获取粉丝列表** `/fans/list/`
```go
list, err := manager.FansList(douyinGo.FansListReq{
    AccessToken: "ACCESS_TOKEN",
    OpenId:      "OPEN_ID",
    Count:       10,
})
```

**获取关注列表** `/following/list/`
```go
list, err := manager.FollowingList(douyinGo.FollowingListReq{
    AccessToken: "ACCESS_TOKEN",
    OpenId:      "OPEN_ID",
    Count:       10,
})
```

**上传视频到文件服务器** `/api/douyin/v1/video/upload_video/`
```go
rs, err := manager.VideoUpload(douyinGo.VideoUploadReq{
    AccessToken: "ACCESS_TOKEN",
    OpenId:      "OPEN_ID",
    FilePath:    "FILE_PATH",
})
```

**分片初始化上传** `/api/douyin/v1/video/init_video_part_upload/`
```go
rs, err := manager.VideoPartUploadInit(douyinGo.VideoPartUploadInitReq{
    AccessToken: "ACCESS_TOKEN",
    OpenId:      "OPEN_ID",
})
```

**分片上传视频** `/api/douyin/v1/video/upload_video_part/`
```go
rs, err := manager.VideoPartUpload(douyinGo.VideoPartUploadReq{
    AccessToken: "ACCESS_TOKEN",
    OpenId:      "OPEN_ID",
    FilePath:    "FILE",
    UploadId:    "UPLOAD_ID",
    ChunkSize:   5 * 1024 * 1024,
    Workers:     4,
})
```

**分片完成上传** `/api/douyin/v1/video/complete_video_part_upload/`
```go
rs, err := manager.VideoUploadPartComplete(douyinGo.VideoUploadPartCompleteReq{
    AccessToken: "ACCESS_TOKEN",
    OpenId:      "OPEN_ID",
    UploadId:    "UPLOAD_ID",
})
```

**创建抖音视频** `/api/douyin/v1/video/create_video/`
```go
rs, err := manager.VideoCreate(douyinGo.VideoCreateReq{
    AccessToken: "ACCESS_TOKEN",
    OpenId:      "OPEN_ID",
    Body: douyinGo.VideoCreateBody{
        VideoId: "VIDEO_ID",
        Text:    "TITLE",
    },
})
```

**删除授权用户发布的视频** `/video/delete/`
```go
rs, err := manager.VideoDelete(douyinGo.VideoDeleteReq{
    AccessToken: "ACCESS_TOKEN",
    OpenId:      "OPEN_ID",
    Body: douyinGo.VideoDeleteBody{
        ItemId: "VIDEO_ID",
    },
})
```

**上传图片到文件服务器** `/api/douyin/v1/video/upload_image/`
```go
rs, err := manager.ImageUpload(douyinGo.ImageUploadReq{
    AccessToken: "ACCESS_TOKEN",
    OpenId:      "OPEN_ID",
    FilePath:    "FILE",
})
```

**发布图片** `/api/douyin/v1/video/create_image_text/`
```go
rs, err := manager.ImageCreate(douyinGo.ImageCreateReq{
    AccessToken: "ACCESS_TOKEN",
    OpenId:      "OPEN_ID",
    Body: douyinGo.ImageCreateBody{
        ImageList: []string{"IMAGE_ID"},
        Text:    "TITLE",
    },
})
```

**查询授权账号视频列表** `/video/list/`
```go
list, err := manager.VideoList(douyinGo.VideoListReq{
    OpenId:      "OPEN_ID",
    AccessToken: "ACCESS_TOKEN",
    Cursor:      0,
    Count:       10,
})
```

**查询指定视频数据** `/video/data/`
```go
rs, err := manager.VideoData(douyinGo.VideoDataReq{
    AccessToken: "ACCESS_TOKEN",
    OpenId:      "OPEN_ID",
    Body: douyinGo.VideoDataBody{
        ItemIds: []string{"VIDEO_ID"},
    },
})
```

**获取share-id** `/share-id/`
```go
rs, err := manager.AwemeShare(douyinGo.AwemeShareReq{
    AccessToken:  "CLIENT_TOKEN",
    NeedCallBack: true,
})
```

**评论列表** `/item/comment/list/`
```go
list, err := manager.ItemCommentList(douyinGo.ItemCommentListReq{
    AccessToken: "ACCESS_TOKEN",
    OpenId:      "OPEN_ID",
    ItemId:      "VIDEO_ID",
    Cursor:      0,
    Count:       10,
})
```

**评论回复列表** `/item/comment/reply/list/`
```go
list, err := manager.ItemCommentReplyList(douyinGo.ItemCommentReplyListReq{
    AccessToken: "ACCESS_TOKEN",
    OpenId:      "OPEN_ID",
    ItemId:      "VIDEO_ID",
    CommentId:   "COMMENT_ID",
    Cursor:      0,
    Count:       10,
})
```

**回复视频评论** `/item/comment/reply/`
```go
rs, err := manager.ItemCommentReply(douyinGo.ItemCommentReplyReq{
    AccessToken: "ACCESS_TOKEN",
    OpenId:      "OPEN_ID",
    Body: douyinGo.ItemCommentReplyBody{
        CommentId: "COMMENT_ID",
        ItemId:    "VIDEO_ID",
        Content:   "CONTENT",
    },
})
```

**置顶视频评论** `/item/comment/top/`
```go
rs, err := manager.ItemCommentTop(douyinGo.ItemCommentTopReq{
    AccessToken: "ACCESS_TOKEN",
    OpenId:      "OPEN_ID",
    Body: douyinGo.ItemCommentTopBody{
        CommentId: "COMMENT_ID",
        ItemId:    "VIDEO_ID",
        Top:       true,
    },
})
```

**获取用户视频情况** `/data/external/user/item/`
```go
rs, err := manager.DataExternalUserItem(douyinGo.DataExternalUserItemReq{
    AccessToken: "ACCESS_TOKEN",
    OpenId:      "OPEN_ID",
    DataType:    30,
})
```

**获取用户粉丝数** `/data/external/user/fans/`
```go
rs, err := manager.DataExternalUserFans(douyinGo.DataExternalUserFansReq{
    AccessToken: "ACCESS_TOKEN",
    OpenId:      "OPEN_ID",
    DataType:    30,
})
```

**获取用户点赞数** `/data/external/user/like/`
```go
rs, err := manager.DataExternalUserLike(douyinGo.DataExternalUserLikeReq{
    AccessToken: "ACCESS_TOKEN",
    OpenId:      "OPEN_ID",
    DataType:    30,
})
```

**获取用户评论数** `/data/external/user/comment/`
```go
rs, err := manager.DataExternalUserComment(douyinGo.DataExternalUserCommentReq{
    AccessToken: "ACCESS_TOKEN",
    OpenId:      "OPEN_ID",
    DataType:    30,
})
```

**获取用户分享数** `/data/external/user/share/`
```go
rs, err := manager.DataExternalUserShare(douyinGo.DataExternalUserShareReq{
    AccessToken: "ACCESS_TOKEN",
    OpenId:      "OPEN_ID",
    DataType:    30,
})
```

**获取用户主页访问数** `/data/external/user/profile/`
```go
rs, err := manager.DataExternalUserProfile(douyinGo.DataExternalUserProfileReq{
    AccessToken: "ACCESS_TOKEN",
    OpenId:      "OPEN_ID",
    DataType:    30,
})
```

**获取视频基础数据** `/data/external/item/base/`
```go
rs, err := manager.DataExternalItemBase(douyinGo.DataExternalItemBaseReq{
    AccessToken: "ACCESS_TOKEN",
    OpenId:      "OPEN_ID",
    ItemId:      "VIDEO_ID",
})
```

**获取视频点赞数据** `/data/external/item/like/`
```go
rs, err := manager.DataExternalItemLike(douyinGo.DataExternalItemLikeReq{
    AccessToken: "ACCESS_TOKEN",
    OpenId:      "OPEN_ID",
    ItemId:      "VIDEO_ID",
    DateType:    7,
})
```

**获取视频评论数据** `/data/external/item/comment/`
```go
rs, err := manager.DataExternalItemComment(douyinGo.DataExternalItemCommentReq{
    AccessToken: "ACCESS_TOKEN",
    OpenId:      "OPEN_ID",
    ItemId:      "VIDEO_ID",
    DateType:    7,
})
```

**获取视频播放数据** `/data/external/item/play/`
```go
rs, err := manager.DataExternalItemPlay(douyinGo.DataExternalItemPlayReq{
    AccessToken: "ACCESS_TOKEN",
    OpenId:      "OPEN_ID",
    ItemId:      "VIDEO_ID",
    DateType:    7,
})
```

**获取视频分享数据** `/data/external/item/share/`
```go
rs, err := manager.DataExternalItemShare(douyinGo.DataExternalItemShareReq{
    AccessToken: "ACCESS_TOKEN",
    OpenId:      "OPEN_ID",
    ItemId:      "VIDEO_ID",
    DateType:    7,
})
```

**获取实时热点词** `/hotsearch/sentences/`
```go
rs, err := manager.HotSearchSentences(douyinGo.HotSearchSentencesReq{
    AccessToken: "CLIENT_TOKEN",
})
```

**获取上升词** `/hotsearch/trending/sentences/`
```go
rs, err := manager.HotSearchTrendingSentences(douyinGo.HotSearchTrendingSentencesReq{
    AccessToken: "CLIENT_TOKEN",
    Count:       10,
})
```

**获取热点词聚合的视频** `/hotsearch/videos/`
```go
rs, err := manager.HotSearchVideos(douyinGo.HotSearchVideosReq{
    AccessToken: "CLIENT_TOKEN",
    HotSentence: "HOT_SENTENCE",
})
```

**获取抖音星图达人热榜** `/star/hot_list/`
```go
list, err := manager.StarHotList(douyinGo.StarHotListReq{
    AccessToken: "CLIENT_TOKEN",
    HotListType: 1,
})
```

**获取抖音星图达人指数** `/star/author_score/`
```go
rs, err := manager.StarAuthorScore(douyinGo.StarAuthorScoreReq{
    AccessToken: "ACCESS_TOKEN",
    OpenId:      "OPEN_ID",
})
```

**获取抖音星图达人指数数据V2** `/star/author_score_v2/`
```go
rs, err := manager.StarAuthorScoreV2(douyinGo.StarAuthorScoreV2Req{
    AccessToken: "CLIENT_TOKEN",
    UniqueId:    "UNIQUE_ID",
})
```

**获取事件订阅状态** `/event/status/list/`
```go
rs, err := manager.EventStatusList(douyinGo.EventStatusListReq{
    AccessToken: "CLIENT_TOKEN",
})
```

**更新应用推送事件订阅状态** `/event/status/update/`
```go
rs, err := manager.EventStatusUpdate(douyinGo.EventStatusUpdateReq{
    AccessToken: "CLIENT_TOKEN",
    Body: douyinGo.EventStatusUpdateBody{
        List: []douyinGo.EventStatus{
            douyinGo.EventStatus{
                Event: "create_video",
                Status: 1,
            },
            douyinGo.EventStatus{
                Event: "authorize",
                Status: 0,
            },
        },
    },
})
```

**消息来源验证**
```go
// 如果使用 github.com/gin-gonic/gin 获取 body和signature进行验证
body, _ := c.GetRawData()
signature := c.GetHeader("X-Douyin-Signature")
manager.WebHookSignature(body, signature)
```

**获取jsapi_ticket** `/js/getticket/`
```go
ticket, err := manager.JsTicket(douyinGo.JsTicketReq{
    AccessToken: "CLIENT_TOKEN",
})
```

**根据jsapi_ticket和其他字段进行签名计算**
```go
signature := manager.JsConfigSignature(douyinGo.ConfigSignReq{
    JsTicket:  "JSAPI_TICKET",
    Timestamp: "TIMESTAMP",
    NonceStr:  "NONCE_STR",
    Url:       "URL",
})
```

**获取用户粉丝数据** `/api/douyin/v1/user/fans_data/`
```go
rs, err := manager.FansData(douyinGo.FansDataReq{
    AccessToken: "ACCESS_TOKEN",
    OpenId:      "OPEN_ID",
})
```

**获取用户粉丝来源分布** `/data/extern/fans/source/`
```go
rs, err := manager.DataExternalFansSource(douyinGo.DataExternalFansSourceReq{
    AccessToken: "ACCESS_TOKEN",
    OpenId:      "OPEN_ID",
})
```

**获取用户粉丝喜好** `/data/extern/fans/favourite/`
```go
rs, err := manager.DataExternalFansFavourite(douyinGo.DataExternalFansFavouriteReq{
    AccessToken: "ACCESS_TOKEN",
    OpenId:      "OPEN_ID",
})
```

**获取用户粉丝热评** `/data/extern/fans/comment/`
```go
rs, err := manager.DataExternalFansComment(douyinGo.DataExternalFansCommentReq{
    AccessToken: "ACCESS_TOKEN",
    OpenId:      "OPEN_ID",
})
```

**获取抖音电影榜、抖音电视剧榜、抖音综艺榜** `/discovery/ent/rank/item/`
```go
rs, err := manager.DiscoveryEntRankItem(douyinGo.DiscoveryEntRankItemReq{
    AccessToken: "CLIENT_TOKEN",
    Type:        1,
})
```

**获取抖音影视综榜单版本** `/discovery/ent/rank/version/`
```go
rs, err := manager.DiscoveryEntRankVersion(douyinGo.DiscoveryEntRankVersionReq{
    AccessToken: "CLIENT_TOKEN",
    Cursor:      0,
    Count:       10,
    Type:        1,
})
```

**获取达人榜单数据** `/data/extern/billboard/stars/`
```go
rs, err := manager.DataExternalBillboard(douyinGo.DataExternalBillboardReq{
    AccessToken: "CLIENT_TOKEN",
    Uri:         douyinGoConf.API_DATA_EXTERNAL_BILLBOARD_STARS, // 响应参数一致,参考 conf/conf.go 替换请求链接
})
```

**获取道具榜单数据** `/data/extern/billboard/prop/`
```go
rs, err := manager.DataExternalBillboardProp(douyinGo.DataExternalBillboardPropReq{
    AccessToken: "CLIENT_TOKEN",
})
```

**获取热门视频数据** `/data/extern/billboard/hot_video/`
```go
rs, err := manager.DataExternalBillboardHotVideo(douyinGo.DataExternalBillboardHotVideoReq{
    AccessToken: "CLIENT_TOKEN",
})
```

**获取直播榜数据** `/data/extern/billboard/live/`
```go
rs, err := manager.DataExternalBillboardLive(douyinGo.DataExternalBillboardLiveReq{
    AccessToken: "CLIENT_TOKEN",
})
```

**获取音乐榜单数据** `/data/extern/billboard/music/hot/`
```go
rs, err := manager.DataExternalBillboardMusic(douyinGo.DataExternalBillboardMusicReq{
    AccessToken: "CLIENT_TOKEN",
    Uri:         douyinGoConf.API_DATA_EXTERNAL_BILLBOARD_MUSIC_HOT, // 响应参数一致,参考 conf/conf.go 替换请求链接
})
```

**查询POI信息** `/poi/search/keyword/`
```go
rs, err := manager.PoiSearchKeyword(douyinGo.PoiSearchKeywordReq{
    AccessToken: "CLIENT_TOKEN",
    Cursor:      0,
    Count:       10,
    Keyword:     "美食",
    City:        "北京",
})
```

**关键词视频搜索** `/video/search/`
```go
rs, err := manager.VideoSearch(douyinGo.VideoSearchReq{
    AccessToken: "ACCESS_TOKEN",
    OpenId:      "OPEN_ID",
    Count:       10,
    Cursor:      0,
    Keyword:     "美食",
})
```

**关键词视频评论列表** `/video/search/comment/list/`
```go
rs, err := manager.VideoSearchCommentList(douyinGo.VideoSearchCommentListReq{
    AccessToken: "CLIENT_TOKEN",
    Count:       10,
    Cursor:      0,
    SecItemId:   "SEC_ITEM_ID",
})
```

**关键词视频评论回复列表** `/video/search/comment/reply/list/`
```go
rs, err := manager.VideoSearchCommentReplyList(douyinGo.VideoSearchCommentReplyListReq{
    AccessToken: "CLIENT_TOKEN",
    Count:       10,
    Cursor:      0,
    SecItemId:   "SEC_ITEM_ID",
    CommentId:   "COMMENT_ID",
})
```

**关键词视频评论回复** `/video/search/comment/reply/`
```go
rs, err := manager.VideoSearchCommentReply(douyinGo.VideoSearchCommentReplyReq{
    AccessToken: "ACCESS_TOKEN",
    OpenId:      "OPEN_ID",
    Body: douyinGo.VideoSearchCommentReplyBody{
        CommentId: "COMMENT_ID",
        SecItemId: "SEC_ITEM_ID",
        Content:   "CONTENT",
    },
})
```

**粉丝判断** `/fans/check/`
```go
rs, err := manager.FansCheck(douyinGo.FansCheckReq{
    AccessToken:    "ACCESS_TOKEN",
    OpenId:         "OPEN_ID",
    FollowerOpenId: "FOLLOWER_OPEN_ID",
})
```

**用户经营身份管理** `/api/douyin/v1/role/check/`
```go
rs, err := manager.FansCheck(douyinGo.FansCheckReq{
    AccessToken:    "CLIENT_TOKEN",
    OpenId:         "OPEN_ID",
    DouyinShortId:  "DOUYIN_SHORTID",
    RoleLabels:     "ROLE_LABELS",
})
```

**通过VideoID获取IFrame代码** `/api/douyin/v1/video/get_iframe_by_video/`
```go
rs, err := manager.VideoGetIframeByVideo(douyinGo.VideoGetIframeByVideoReq{
    VideoId:    "VIDEO_ID",
})
```

**通过ItemID获取IFrame代码** `/api/douyin/v1/video/get_iframe_by_item/`
```go
rs, err := manager.VideoGetIframeByItem(douyinGo.VideoGetIframeByItemReq{
    ClientKey: "CLIENT_KEY",
    ItemId:    "ITEM_ID",
})
```

**获取open_ticket** `/open/getticket/`
```go
rs, err := manager.OpenTicket(douyinGo.OpenTicketReq{
    AccessToken: "CLIENT_TOKEN",
})
```

**H5分享跳转链接获取** `/api/douyin/v1/schema/get_share/`
```go
rs, err := manager.SchemaGetShare(douyinGo.SchemaGetShareReq{
    AccessToken: "CLIENT_TOKEN",
    Body: douyinGo.SchemaGetShareBody{
        ClientTicket: "OPEN_TICKET",
        ExpireAt:     "EXPIRE_AT",
        VideoPath:    "VIDEO_PATH",
    },
})
```

**个人页跳转链接获取** `/api/douyin/v1/schema/get_user_profile/`
```go
rs, err := manager.SchemaGetUserProfile(douyinGo.SchemaGetUserProfileReq{
    AccessToken: "CLIENT_TOKEN",
    Body: douyinGo.SchemaGetUserProfileBody{
        ExpireAt:     "EXPIRE_AT",
        OpenId:       "OPEN_ID",
    },
})
```

**个人会话页跳转链接获取** `/api/douyin/v1/schema/get_chat/`
```go
rs, err := manager.SchemaGetChat(douyinGo.SchemaGetChatReq{
    AccessToken: "CLIENT_TOKEN",
    Body: douyinGo.SchemaGetChatBody{
        ExpireAt:     "EXPIRE_AT",
        OpenId:       "OPEN_ID",
    },
})
```

**视频详情页跳转链接获取** `/api/douyin/v1/schema/get_item_info/`
```go
rs, err := manager.SchemaGetItemInfo(douyinGo.SchemaGetItemInfoReq{
    AccessToken: "CLIENT_TOKEN",
    Body: douyinGo.SchemaGetItemInfoBody{
        ExpireAt:     "EXPIRE_AT",
        ItemId:       "ITEM_ID",
    },
})
```

**直播间跳转链接获取** `/api/douyin/v1/schema/get_live/`
```go
rs, err := manager.SchemaGetLive(douyinGo.SchemaGetLiveReq{
    AccessToken: "CLIENT_TOKEN",
    Body: douyinGo.SchemaGetLiveBody{
        ExpireAt:     "EXPIRE_AT",
        OpenId:       "OPEN_ID",
    },
})
```