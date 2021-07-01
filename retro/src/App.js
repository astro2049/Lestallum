// https://cloud.tencent.com/document/product/436/35649

import { makeStyles } from "@material-ui/core/styles";
import { DropzoneArea } from "material-ui-dropzone";
import Button from "@material-ui/core/Button";
import { useState } from "react";

const { REACT_APP_SECRET_ID, REACT_APP_SECRET_KEY } = process.env;

var COS = require("cos-js-sdk-v5");
var cos = new COS({
    SecretId: REACT_APP_SECRET_ID,
    SecretKey: REACT_APP_SECRET_KEY,
});

const useStyles = makeStyles(() => ({
    outerContainer: {
        width: "100vw",
        height: "100vh",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
    },
    fileDropzone: {
        width: "600px",
        height: "300px",
        display: "flex",
        flexDirection: "column",
    },
}));

function App() {
    const classes = useStyles();

    const [file, setFile] = useState();

    const handleSave = (files) => {
        console.log(files);
        setFile(files[0]);
        console.log(files[0]);
    };

    const uploadFile = () => {
        cos.uploadFile(
            {
                Bucket: "paradise-1305422781" /* 必须 */,
                Region: "ap-beijing" /* 存储桶所在地域，必须字段 */,
                Key: file.path /* 必须 */,
                Body: file /* 必须 */,
                SliceSize:
                    1024 *
                    1024 *
                    5 /* 触发分块上传的阈值，超过5MB使用分块上传，非必须 */,
                onTaskReady: function (taskId) {
                    /* 非必须 */
                    console.log(taskId);
                },
                onProgress: function (progressData) {
                    /* 非必须 */
                    console.log(JSON.stringify(progressData));
                },
                onFileFinish: function (err, data, options) {
                    console.log(options.Key + "上传" + (err ? "失败" : "完成"));
                },
            },
            function (err, data) {
                console.log(err || data);
            }
        );
    };

    return (
        <div className={classes.outerContainer}>
            <div className={classes.fileDropzone}>
                <DropzoneArea filesLimit="1" onChange={(e) => handleSave(e)} />
                <Button
                    variant="contained"
                    color="primary"
                    onClick={() => uploadFile()}
                >
                    Submit
                </Button>
            </div>
        </div>
    );
}

export default App;
