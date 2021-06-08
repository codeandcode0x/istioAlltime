package com.msa.ticket.frontend.pojo;

import com.msa.ticket.frontend.model.Info;

import java.util.List;

public class InfoDatas {
    private Integer code;
    private List<Info> data;

    public Integer getCode() {
        return code;
    }

    public void setCode(Integer code) {
        this.code = code;
    }

    public List<Info> getData() {
        return data;
    }

    public void setData(List<Info> data) {
        this.data = data;
    }

    @Override
    public String toString() {
        return "Message{" +
                "code='" + code + '\'' +
                ", data=" + data +
                '}';
    }
}


