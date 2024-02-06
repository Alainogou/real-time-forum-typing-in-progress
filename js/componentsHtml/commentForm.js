function renderCommentForm(container,  postId) {
    // Rendu du formulaire de commentaire
    container.innerHTML = `
        <hr>
        <div><span class="errorStyle EmptyContent"></span></div> 
        <div class="comment_warpper">
       
        
            <div class="user" style="background-color:#efefef; height:30px;width:30px; text-align:center; border-radius:50%; padding-top:4px">
                <i class="fa-solid fa-user" ></i>
            </div>
           
            <div class="comment_search">

                <form  enctype="multipart/form-data" class="commentform-${postId}  sendComent" >
                        <input type="hidden" name="post_id" value="${postId}">

                        <input class="nc-ct" type="text" name="content" placeholder= "write your comment here...">
                      
                         
                        <button>
                                    <div class="svg-wrapper-1">
                                        <div class="svg-wrapper">
                                        <svg
                                            xmlns="http://www.w3.org/2000/svg"
                                            viewBox="0 0 24 24"
                                            width="24"
                                            height="24"
                                        >
                                            <path fill="none" d="M0 0h24v24H0z"></path>
                                            <path
                                            fill="currentColor"
                                            d="M1.946 9.315c-.522-.174-.527-.455.01-.634l19.087-6.362c.529-.176.832.12.684.638l-5.454 19.086c-.15.529-.455.547-.679.045L12 14l6-8-8 6-8.054-2.685z"
                                            ></path>
                                        </svg>
                                        </div>
                                    </div>
                                    <span>Send</span>
                        </button>

                        
                </form>
            </div>

           
      
        </div>
    `;

   
}

export {renderCommentForm}


// {/* <div class="comment_search">
// <input type="text" name="commentText" placeholder="Write a comment">
// <input type="hidden" name="user_commented" value="${userId}">

// <i class="fa-solid fa-paper-plane"></i>  
        
// </div> */}