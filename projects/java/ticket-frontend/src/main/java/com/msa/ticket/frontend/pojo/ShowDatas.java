package com.msa.ticket.frontend.pojo;

import com.msa.ticket.frontend.model.Show;

import java.util.List;

public class ShowDatas {
    private Integer code;
    private List<Show> data;

    public Integer getCode() {
        return code;
    }

    public void setCode(Integer code) {
        this.code = code;
    }

    public List<Show> getData() {
        return data;
    }

    public void setData(List<Show> data) {
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


