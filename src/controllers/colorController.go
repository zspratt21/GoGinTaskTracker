package controllers

import (
	"GoDynamoApiTemplate/src/database"
	"GoDynamoApiTemplate/src/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gin-gonic/gin"
)

func AddColor(c *gin.Context) {
	var color models.Color

	if err := c.ShouldBindJSON(&color); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request format"})
		return
	}

	sess, err := database.NewSession()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create AWS session"})
		return
	}
	svc := dynamodb.New(sess)

	av, err := dynamodbattribute.MarshalMap(color)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to marshal color"})
		return
	}

	_, err = svc.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String("colors"),
		Item:      av,
	})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to add color", "message": err.Error()})
		return
	}

	c.JSON(201, color)
}

func DeleteColor(c *gin.Context) {
	colorName := c.Param("name")

	sess, err := database.NewSession()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create AWS session"})
		return
	}
	svc := dynamodb.New(sess)

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("colors"),
		Key: map[string]*dynamodb.AttributeValue{
			"name": {
				S: aws.String(colorName),
			},
		},
	})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch color for deletion"})
		return
	}
	if result.Item == nil {
		c.JSON(404, gin.H{"error": "Color not found"})
		return
	}

	_, err = svc.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String("colors"),
		Key: map[string]*dynamodb.AttributeValue{
			"name": {
				S: aws.String(colorName),
			},
		},
	})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete color", "details": err.Error()})
		return
	}

	c.JSON(200, gin.H{"success": "Color deleted"})
}

func GetColor(c *gin.Context) {
	colorName := c.Param("name")

	sess, err := database.NewSession()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create AWS session"})
		return
	}
	svc := dynamodb.New(sess)

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("colors"),
		Key: map[string]*dynamodb.AttributeValue{
			"name": {
				S: aws.String(colorName),
			},
		},
	})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch color"})
		return
	}
	if result.Item == nil {
		c.JSON(404, gin.H{"error": "Color not found"})
		return
	}

	var color models.Color
	err = dynamodbattribute.UnmarshalMap(result.Item, &color)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to unmarshal color"})
		return
	}

	c.JSON(200, color)
}

func GetColors(c *gin.Context) {
	sess, err := database.NewSession()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create AWS session"})
		return
	}
	svc := dynamodb.New(sess)

	result, err := svc.Scan(&dynamodb.ScanInput{
		TableName: aws.String("colors"),
	})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch colors"})
		return
	}

	var colors []models.Color
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &colors)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to unmarshal colors"})
		return
	}

	c.JSON(200, colors)
}

func UpdateColor(c *gin.Context) {
	var color models.Color

	if err := c.ShouldBindJSON(&color); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request format"})
		return
	}

	colorName := c.Param("name")

	if colorName != color.Name {
		c.JSON(400, gin.H{"error": "Color name in URL does not match name in request body"})
		return
	}

	sess, err := database.NewSession()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create AWS session"})
		return
	}
	svc := dynamodb.New(sess)

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("colors"),
		Key: map[string]*dynamodb.AttributeValue{
			"name": {
				S: aws.String(colorName),
			},
		},
	})
	if err != nil {
		c.JSON(500, gin.H{"error": "Error checking if color exists"})
		return
	}

	if result.Item == nil {
		c.JSON(404, gin.H{"error": "Color not found"})
		return
	}

	av, err := dynamodbattribute.MarshalMap(color)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to marshal color"})
		return
	}

	_, err = svc.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String("colors"),
		Item:      av,
	})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update color", "details": err.Error()})
		return
	}

	c.JSON(200, color)
}
