# douyin-go
抖音Open API
```
// 初始化
credentials := douyinGo.NewCredentials("CLIENT_KEY", "CLIENT_SECRET")
manager := douyinGo.NewManager(credentials, nil)

// 生成授权链接
oauthUrl := manager.OauthConnect(douyinGo.OauthParam{
  Scope: "user_info,mobile_alert,video.list,video.data,video.create,video.delete,data.external.user,data.external.item,aweme.share,fans.list,following.list,item.comment,star_top_score_display",
  RedirectUri: "REDIRECT_URI",
})

// 获取AccessToken
accessToken, err := manager.OauthAccessToken(douyinGo.OauthAccessTokenReq{
  Code: "CODE",
})

// 获取视频列表
list, err := manager.VideoList(douyinGo.VideoListReq{
  OpenId:      "OPEN_ID",
  AccessToken: "ACCESS_TOKEN",
  Cursor:      0,
  Count:       10,
})
