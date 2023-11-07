package controllers

import (
	"GoDynamoApiTemplate/src/database"
	"GoDynamoApiTemplate/src/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AddTask(c *gin.Context) {
	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request format"})
		return
	}

	task.Id = uuid.New().String()

	sess, err := database.NewSession()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create AWS session"})
		return
	}
	svc := dynamodb.New(sess)

	av, err := dynamodbattribute.MarshalMap(task)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to marshal task"})
		return
	}

	_, err = svc.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String("tasks"),
		Item:      av,
	})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to add task", "message": err.Error()})
		return
	}

	c.JSON(201, task)
}

func DeleteTask(c *gin.Context) {
	taskId := c.Param("id")

	sess, err := database.NewSession()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create AWS session"})
		return
	}
	svc := dynamodb.New(sess)

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("tasks"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(taskId),
			},
		},
	})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch task for deletion"})
		return
	}
	if result.Item == nil {
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}

	_, err = svc.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String("tasks"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(taskId),
			},
		},
	})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete task", "details": err.Error()})
		return
	}

	c.JSON(200, gin.H{"success": "Task deleted"})
}

func GetTask(c *gin.Context) {
	taskId := c.Param("id")

	sess, err := database.NewSession()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create AWS session"})
		return
	}
	svc := dynamodb.New(sess)

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("tasks"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(taskId),
			},
		},
	})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch task"})
		return
	}
	if result.Item == nil {
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}

	var task models.Task
	err = dynamodbattribute.UnmarshalMap(result.Item, &task)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to unmarshal task"})
		return
	}

	c.JSON(200, task)
}

func GetTasks(c *gin.Context) {
	sess, err := database.NewSession()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create AWS session"})
		return
	}
	svc := dynamodb.New(sess)

	result, err := svc.Scan(&dynamodb.ScanInput{
		TableName: aws.String("tasks"),
	})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch tasks"})
		return
	}

	var tasks []models.Task
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &tasks)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to unmarshal tasks"})
		return
	}

	c.JSON(200, tasks)
}

func UpdateTask(c *gin.Context) {
	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request format"})
		return
	}

	taskId := c.Param("id")

	if taskId != task.Id {
		c.JSON(400, gin.H{"error": "task id in URL does not match id in request body"})
		return
	}

	sess, err := database.NewSession()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create AWS session"})
		return
	}
	svc := dynamodb.New(sess)

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("tasks"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(taskId),
			},
		},
	})
	if err != nil {
		c.JSON(500, gin.H{"error": "Error checking if task exists"})
		return
	}

	if result.Item == nil {
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}

	av, err := dynamodbattribute.MarshalMap(task)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to marshal task"})
		return
	}

	_, err = svc.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String("tasks"),
		Item:      av,
	})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update task", "details": err.Error()})
		return
	}

	c.JSON(200, task)
}
