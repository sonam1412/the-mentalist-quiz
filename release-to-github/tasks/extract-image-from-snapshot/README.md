# extract-image-from-snapshot

Tekton task that extracts the container image to be released on github.com from the snapshot.

The URL of the container image built by Konflux, is provided as a result.

## Parameters

| Name                | Description                                                                | Optional | Default value |
|---------------------|----------------------------------------------------------------------------|----------|---------------|
| snapshotPath        | Path to the JSON string of the mapped Snapshot spec in the data workspace  | No       | -             |
| subdirectory        | Subdirectory inside the workspace to be used for storing the snapshotSpec  | Yes      | ""            |

## Results

| Name      | Description                           |
|-----------|---------------------------------------|
| image_url | The URL of the image built by Konflux |
