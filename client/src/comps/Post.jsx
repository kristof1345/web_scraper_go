import React from "react";

const Post = ({ post }) => {
  console.log(post.ID);
  return (
    <div className="post">
      <div className="postID">{post.ID}</div>
      <div className="postDesc">{post.Description}</div>
    </div>
  );
};

export default Post;
