import React, { useEffect, useState } from "react"
import detailcss from "../styles/detail.module.css"
import testdata_satu from "../testdata/testdata_satu"

const detail = () => {
    // const URL = 'http://localhost:3030/clickDetail/getDetail'
    const URL_PROD = 'http://bacotsantuy.ddns.net:80/clickDetail/getDetail'
    // const URL_DEV = 'http://www.json-generator.com/api/json/get/craPDFlpxK?indent=2'

    const GetDetail = () => {
        const [detaildata, setDetail] = useState([])
        const [loading, setLoading] = useState(false)

        useEffect(() => {
            try {
                async function fetchdata() {
                    const rawdata = await fetch(URL_PROD)
                    const jsondata = await rawdata.json()
                    setDetail(jsondata)
                    setLoading(true)
                }
                fetchdata()
            } catch (error) {
                
            }
        },[])
        if (loading === true) {
            return (
                <div className={detailcss.detail_ctn}>
                    <div className={detailcss.detail_title}>{detaildata.ContentTitle}</div>
                    <div>{detaildata.ContentContent}</div>
                    <div className={detailcss.detail_tag} style={{background: detaildata.Tag.TagColor}}>{detaildata.Tag.TagName}</div>
                </div>
            )
        } else {
            return (
                <div className={detailcss.detail_loading}>Getting detail..</div>
            )
        }
    }

    return (
        <div className={detailcss.detail_main}>
            <a className={detailcss.detail_back} href="/">Back</a>
            {GetDetail()}
        </div>
    )
}

export default detail