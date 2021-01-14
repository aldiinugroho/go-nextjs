import React from "react"
import detailcss from "../styles/detail.module.css"
import testdata_satu from "../testdata/testdata_satu"

const detail = () => {

    const GetDetail = () => {
        const URL = testdata_satu[0]
        return (
            <div className={detailcss.detail_ctn}>
                <div className={detailcss.detail_title}>{URL.title}</div>
                <div>{URL.content}</div>
            </div>
        )
    }

    return (
        <div className={detailcss.detail_main}>
            <a className={detailcss.detail_back} href="/">Back</a>
            {GetDetail()}
        </div>
    )
}

export default detail