import React from "react";

const Post = ({ post }) => {
  return (
    <a className="post" href={post.URL} target="_blank">
      <div className="postDesc">{post.Description}</div>
    </a>
  );
};

export default Post;
