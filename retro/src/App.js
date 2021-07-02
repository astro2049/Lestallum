import { makeStyles } from "@material-ui/core/styles";
import { DropzoneArea } from "material-ui-dropzone";
import Button from "@material-ui/core/Button";
import { useState } from "react";
import axios from "axios";

const { REACT_APP_SERVER_ADDRESS } = process.env;

var COS = require("cos-js-sdk-v5");

// https://cloud.tencent.com/document/product/436/11459
var cos = new COS({
    getAuthorization: function (options, callback) {
        // 异步获取临时密钥
        axios.get(REACT_APP_SERVER_ADDRESS + "/file/sts").then((response) => {
            var data = response.data;
            var credentials = data.Credentials;
            if (!data || !credentials) {
                return console.error("credentials invalid");
            }
            callback({
                TmpSecretId: credentials.TmpSecretId,
                TmpSecretKey: credentials.TmpSecretKey,
                SecurityToken: credentials.Token,
                // 建议返回服务器时间作为签名的开始时间，避免用户浏览器本地时间偏差过大导致签名错误
                StartTime: data.StartTime, // 时间戳，单位秒，如：1580000000
                ExpiredTime: data.ExpiredTime, // 时间戳，单位秒，如：1580000900
            });
        });
    },
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

    const Bucket = "heaven-1305422781";
    const Region = "ap-chengdu";

    const [file, setFile] = useState();

    const handleSave = (files) => {
        setFile(files[0]);
        console.log(files[0]);
    };

    // https://cloud.tencent.com/document/product/436/35649
    const uploadFile = () => {
        cos.uploadFile(
            {
                Bucket: Bucket /* 必须 */,
                Region: Region /* 存储桶所在地域，必须字段 */,
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
                <DropzoneArea
                    filesLimit={1}
                    maxFileSize={1024 * 1024 * 1024 * 10}
                    onChange={(e) => handleSave(e)}
                />
                <Button
                    variant="contained"
                    color="primary"
                    onClick={() => uploadFile()}
                    disabled={file === undefined}
                >
                    Submit
                </Button>
            </div>
        </div>
    );
}

export default App;
