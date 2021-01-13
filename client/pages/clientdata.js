import React, { useEffect, useState } from "react";
const clientdata = () => {
    const URL = 'http://localhost:3030/data'
    const GetResult = () => {
        const [result, setResult] = useState([])

        useEffect(() => {
            try {
                async function fetchData() {
                    const callResult = await fetch(URL)
                    const resultJSN = await callResult.json()
                    setResult(resultJSN)
                }
                fetchData()
                
            } catch (error) {
                
            }
        },[])
        return (
            <div>
                {result.map((itm,key) => (
                    <div key={key}>
                        <div>{itm.Name}</div>
                        <div>{itm.Age}</div>
                    </div>
                ))}
            </div>
        )
    }

    return (
        <div>
            {GetResult()}
        </div>
    )
}

export default clientdata