import React from "react";

const Post = ({ post }) => {
  return (
    <div className="post">
      <div className="postDesc">{post.Description}</div>
    </div>
  );
};

export default Post;
