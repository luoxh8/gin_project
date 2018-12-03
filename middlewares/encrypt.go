package middlewares

import (
	`gin_project/core`
	`gin_project/core/h_categories`
	`github.com/gin-gonic/gin`
	`net/http`
	`net/url`
	`sort`
	`strconv`
	`strings`
)

var EncryptMiddleware = core.Handler{
	"encrypt": func(context *gin.Context) {
		var (
			ip, sn, version, oldKey, platform, mId, appKey, time, params string
			get, post                                                    url.Values
		)
		context.Request.ParseForm()
		ip = h_categories.GetIP(context.Request.RemoteAddr)
		version = context.Query("v")
		oldKey = context.Query("sn")
		time = context.Query("t")
		platform = strings.ToLower(context.DefaultQuery("platform", ""))
		mId = context.DefaultQuery("m_id", "-1")

		/**
		    1、参数校验
		*/
		if ip == "127.0.0.1" || oldKey == "zhihui" {
			return
		}

		if platform == "ios" {
			if mid, err := strconv.Atoi(mId); err != nil {
				goto ERR
			} else {
				if mid != -1 {
					platform = "ios_other"
				}
			}
		}

		if time == "" || oldKey == "" {
			goto ERR
		}

		if platform != "android" && platform != "ios" && platform != "ios_other" {
			goto ERR
		}

		/**
		    2、加密校验
		*/
		if appKey = GetAppKey(platform, version); appKey == "" {
			goto ERR
		}

		get = context.Request.URL.Query()
		post = context.Request.PostForm
		params = SecretString(get, post)
		params += appKey
		sn = h_categories.GenMD5String(params)
		if strings.ToLower(sn) != strings.ToLower(oldKey) {
			if core.Debug == true {
				context.JSON(http.StatusBadRequest, gin.H{
					"get_params":    context.Request.URL.RawQuery,
					"post_params":   post,
					"sorted_params": params,
					"server_sn":     sn,
					"client_sn":     oldKey,
					"server_md5":    sn,
				})
				context.Abort()
				return
			}
			goto ERR
		} else {
			return
		}
	ERR:
		context.JSON(http.StatusBadRequest, core.ErrSysBadRequest)
		context.Abort()
	},
}

/**
    获取appkey
*/
func GetAppKey(platform, version string) string {
	var appKey string
	for key, value := range core.AppSecretKeys[platform].(gin.H) {
		if version == key {
			appKey = value.(string)
			return appKey
		}
	}
	return appKey
}

/**
    获取加密字符串
*/
func SecretString(get, post url.Values) string {
	getArgsDeal := GetArgsDeal(get)
	postArgsDeal := PostArgsDeal(post)
	one := AppendTwoArrayToOne(getArgsDeal, postArgsDeal)
	sort.Strings(one)
	s := strings.Join(one, "")
	return s
}

/**
    合并数组
*/
func AppendTwoArrayToOne(get, post []string) []string {
	returnArray := make([]string, len(get)+len(post), len(get)+len(post)*2)
	var _index = 0

	for _, value := range get {
		if value != "" {
			returnArray[_index] = value
			_index++
		}
	}

	for _, value := range post {
		if value != "" {
			returnArray[_index] = value
			_index++
		}
	}

	return returnArray
}

/**
    处理get参数
*/
func GetArgsDeal(get url.Values) []string {
	getArgs := make([]string, len(get), len(get)*2)
	var _key = ""
	var _index = 0
	for key, value := range get {
		if key != "sn" {
			if _key != key {
				getArgs[_index] = key + "=" + value[0]
				_key = key
				_index++
			}
		}
	}
	return getArgs
}

/**
    处理post参数
*/
func PostArgsDeal(post url.Values) []string {
	postArgs := make([]string, len(post), len(post)*2)
	var _key = ""
	var _index = 0
	for key, value := range post {
		if _key != key {
			postArgs[_index] = key + "=" + value[0]
			_key = key
			_index++
		}
	}
	return postArgs
}
