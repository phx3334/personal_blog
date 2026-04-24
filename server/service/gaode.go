package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"go_blog/server/global"
	"go_blog/server/model/other"
	"go_blog/server/utils"
)

// JwtService 提供与高德相关的服务
type GaodeService struct {
}

// GetLocationByIP 根据 IP 地址获取地理位置信息
// 参数：
//   - ip: 要查询的 IP 地址字符串
// 返回值：
//   - other.IPResponse: 包含 IP 地址对应地理位置信息的响应结构体
//   - error: 错误信息，如 HTTP 请求失败、JSON 解析错误等
func (gaodeService *GaodeService) GetLocationByIP(ip string) (other.IPResponse, error) {
    // 初始化返回结构体
    data := other.IPResponse{}
    // 从全局配置中获取高德地图 API 密钥
    key := global.Config.Gaode.Key
    // 高德地图 IP 定位 API 地址
    urlStr := "https://restapi.amap.com/v3/ip"
    // 请求方法
    method := "GET"
    // 构建请求参数
    params := map[string]string{
        "ip":  ip,    // 要查询的 IP 地址
        "key": key,   // 高德地图 API 密钥
    }
    // 发送 HTTP 请求
    res, err := utils.HttpRequest(urlStr, method, nil, params, nil)
    if err != nil {
        return data, err
    }
    // 延迟关闭响应体
    defer res.Body.Close()

    // 检查 HTTP 响应状态码
    if res.StatusCode != http.StatusOK {
        return data, fmt.Errorf("request failed with status code: %d", res.StatusCode)
    }

    // 读取响应体数据
    byteData, err := io.ReadAll(res.Body)
    if err != nil {
        return data, err
    }

    // 解析 JSON 响应数据到结构体
    err = json.Unmarshal(byteData, &data)
    if err != nil {
        return data, err
    }
    // 返回解析后的数据
    return data, nil
}

// GetWeatherByAdcode 根据城市编码获取实时天气信息
// 参数：
//   - adcode: 城市编码（高德地图标准）
// 返回值：
//   - other.Live: 包含实时天气信息的结构体
//   - error: 错误信息，如 HTTP 请求失败、JSON 解析错误、无天气数据等
func (gaodeService *GaodeService) GetWeatherByAdcode(adcode string) (other.Live, error) {
    // 初始化返回结构体
    data := other.WeatherResponse{}
    // 从全局配置中获取高德地图 API 密钥
    key := global.Config.Gaode.Key
    // 高德地图天气 API 地址
    urlStr := "https://restapi.amap.com/v3/weather/weatherInfo"
    // 请求方法
    method := "GET"
    // 构建请求参数
    params := map[string]string{
        "city": adcode, // 城市编码
        "key":  key,    // 高德地图 API 密钥
    }
    // 发送 HTTP 请求
    res, err := utils.HttpRequest(urlStr, method, nil, params, nil)
    if err != nil {
        return other.Live{}, err
    }
    // 延迟关闭响应体
    defer res.Body.Close()

    // 检查 HTTP 响应状态码
    if res.StatusCode != http.StatusOK {
        return other.Live{}, fmt.Errorf("request failed with status code: %d", res.StatusCode)
    }

    // 读取响应体数据
    byteData, err := io.ReadAll(res.Body)
    if err != nil {
        return other.Live{}, err
    }

    // 解析 JSON 响应数据到结构体
    err = json.Unmarshal(byteData, &data)
    if err != nil {
        return other.Live{}, err
    }

    // 检查是否有返回的天气数据
    if len(data.Lives) == 0 {
        return other.Live{}, fmt.Errorf("no live weather data available") // 没有天气数据时返回错误
    }

    // 返回当天的天气数据
    return data.Lives[0], nil
}