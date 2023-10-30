package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"math"
)

type GeoApi struct {
	redisClient *redis.Client
}

const (
	UserLocationKey = "user_location"
)

// NewGeoApi 创建 GeoApi 实例
func NewGeoApi(redisClient *redis.Client) *GeoApi {
	return &GeoApi{
		redisClient: redisClient,
	}
}

// SaveUserLocation 存自己位置
func (geoApi *GeoApi) SaveUserLocation(ctx context.Context, userId string, longitude, latitude float64) {
	geoApi.redisClient.GeoAdd(ctx, UserLocationKey, &redis.GeoLocation{
		Longitude: longitude,
		Latitude:  latitude,
		Name:      userId,
	})
}

// GetUserLocation 拿自己位置
func (geoApi *GeoApi) GetUserLocation(ctx context.Context, userId string) (longitude, latitude float64) {
	locations, _ := geoApi.redisClient.GeoPos(ctx, UserLocationKey, userId).Result()
	return Round(locations[0].Longitude, 6), Round(locations[0].Latitude, 6)
}

// FindNearestUser 找我最近的人
func (geoApi *GeoApi) FindNearestUser(ctx context.Context, userId string) []redis.GeoLocation {
	// 自己的坐标
	userLocation, _ := geoApi.redisClient.GeoPos(ctx, UserLocationKey, userId).Result()
	// 距离自己最近的user
	nearestUsers, _ := geoApi.redisClient.GeoRadius(ctx,
		UserLocationKey,
		userLocation[0].Longitude,
		userLocation[0].Latitude,
		&redis.GeoRadiusQuery{
			Radius:    100,
			Unit:      "m",
			WithCoord: true,
			WithDist:  true,
			Sort:      "ASC",
			Count:     3,
		}).Result()
	return nearestUsers
}

// GetUserDistance 两个人的距离
func (geoApi *GeoApi) GetUserDistance(ctx context.Context, userId1, userId2 string) float64 {
	distance, _ := geoApi.redisClient.GeoDist(ctx, UserLocationKey, userId1, userId2, "km").Result()
	return Round(distance, 6)
}

// Round 对 float64 值保留指定小数位数
func Round(value float64, places int) float64 {
	multiplier := math.Pow(10, float64(places))
	return math.Round(value*multiplier) / multiplier
}

func main() {
	// 创建Redis客户端
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // 你的Redis服务器地址
	})

	// 创建GeoApi实例
	geoApi := NewGeoApi(redisClient)

	// 调用GeoApi方法
	ctx := context.Background()
	//userId := "456"
	//longitude := 116.58
	//latitude := 33.38

	// 存储用户位置
	//geoApi.SaveUserLocation(ctx, userId, longitude, latitude)

	// 获取用户位置
	//longitude, latitude = geoApi.GetUserLocation(ctx, userId)
	//fmt.Printf("User Location - Longitude: %f, Latitude: %f\n", longitude, latitude)
	//
	//// 找到最近的用户
	//nearestUsers := geoApi.FindNearestUser(ctx, userId)
	//fmt.Printf("Nearest Users: %+v\n", nearestUsers)
	//
	// 获取两个用户之间的距离
	userId1 := "123"
	userId2 := "456"
	distance := geoApi.GetUserDistance(ctx, userId1, userId2)
	fmt.Printf("Distance between %s and %s: %f kilometers\n", userId1, userId2, distance)

}
