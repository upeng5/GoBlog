import React from "react";
import Fab from "@material-ui/core/Fab";
import AddIcon from "@material-ui/icons/AddCircleRounded";

function Feed() {
  return (
    <div>
      {/* Place fab on bottom right side to add a post to feed */}
      <Fab color="primary" aria-label="add">
        <AddIcon />
      </Fab>
    </div>
  );
}

export default Feed;
