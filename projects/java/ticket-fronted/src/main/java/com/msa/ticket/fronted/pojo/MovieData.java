package com.msa.ticket.fronted.pojo;

import com.msa.ticket.fronted.model.Movie;

public class MovieData {
    private Integer code;
    private Movie data;

    public Integer getCode() {
        return code;
    }

    public void setCode(Integer code) {
        this.code = code;
    }

    public Movie getData() {
        return data;
    }

    public void setData(Movie data) {
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


