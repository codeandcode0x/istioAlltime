package com.msa.ticket.frontend.pojo;

import com.msa.ticket.frontend.model.Movie;

import java.io.Serializable;
import java.util.List;

public class MovieDatas implements Serializable {
    private Integer code;
    private List<Movie> data;

    public Integer getCode() {
        return code;
    }

    public void setCode(Integer code) {
        this.code = code;
    }

    public List<Movie> getData() {
        return data;
    }

    public void setData(List<Movie> data) {
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


