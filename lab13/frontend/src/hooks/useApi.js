import React, {createContext, useContext} from "react";

import axios from 'axios'

import {useAuth} from "./useAuth";

const apiContext = createContext();

export function ProvideApi({children}) {
    const auth = useProvideApi();
    return <apiContext.Provider value={auth}>{children}</apiContext.Provider>;
}

export const useApi = () => {
    return useContext(apiContext);
};

function useProvideApi() {
    const auth = useAuth();

    const apiCall = async (method,route,baseUrl="",headers,data={}) => {
        try {
            const response = await axios({
                method: method,
                url: baseUrl + route,
                headers: headers,
                data: data
            });
            const responseData = await response.data;
            return responseData;
        } catch (e) {
            console.log(e)
            return null
        }
    }

    const authApiCall = async (method,route,baseUrl="",headers={},data={}) => {
        const accessToken = await auth.getAccessToken();
        const authHeaders = {Authorization: `Bearer ${accessToken}`}
        return await apiCall(method,route,baseUrl,{...headers,...authHeaders},data)
    }

    return {
        apiCall,
        authApiCall
    };
}