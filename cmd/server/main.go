// Go package
package main

import (
	"net/http"
    "github.com/gin-gonic/gin"
    "errors"
    // "strconv" 
)


type post struct{
    ID string `json:"id"`
    Title string `json:"title"`
    Photo string `json:"photo"`
    Content string `json:"content"`
}

var posts=[]post{
    {
        ID: "1",
        Title :"title",
        Photo :"photo",
        Content :"content",
    },
    {
        ID: "2",
        Title :"title2",
        Photo :"photo2",
        Content :"content2",
    },
    {
        ID: "3",
        Title :"title3",
        Photo :"photo3",
        Content :"content3",
    },
}

//get all posts
func getPosts(context *gin.Context)  {

    context.IndentedJSON(http.StatusOK,posts)
}

//get a post by id
func getPostByID(id string)  (*post, error){
  
    for i, t :=  range posts{
        if t.ID==id{
            return &posts[i],nil
        }
    }
    return nil, errors.New("post not found")
}

//get a post by id
func deletePostByID(id string)  (*post, error){
  
    for i, t :=  range posts{
        if t.ID==id{
         
            // posts=append(posts[:i],posts[i+1:] )
            return &posts[i], nil
        }
    }
    return nil, errors.New("post not found")
}

//get a post
func getPost(context *gin.Context)  {
    id:= context.Params.ByName("postID")
    post,err:= getPostByID((id))
    if err!=nil{
        context.IndentedJSON(http.StatusNotFound,gin.H{
            "message":"Post not found",
        })
        return
    }
    context.IndentedJSON(http.StatusOK,post)
}

//delete a post
func deletePost(context *gin.Context)  {
    id:= context.Params.ByName("postID")
    post,err:= deletePostByID((id))
    if err!=nil{
        context.IndentedJSON(http.StatusNotFound,gin.H{
            "message":"Post not found",
        })
        return
    }
    context.IndentedJSON(http.StatusOK,post)
}


//add a post
func addPost(context *gin.Context)  {
    var newPost post
    if err:= context.BindJSON((&newPost)); err!=nil{
        return
    }

    posts=append(posts, newPost)
    context.IndentedJSON(http.StatusCreated,newPost)
}

func main(){
    router := gin.Default()//server

    router.GET("/posts",getPosts)

    router.GET("/posts/:postID",getPost)
    router.DELETE("/posts/:postID",deletePost)
    router.PATCH("/posts/:postID",deletePost)
    
    router.POST("/post",addPost)

    router.Run("localhost:9090")
}