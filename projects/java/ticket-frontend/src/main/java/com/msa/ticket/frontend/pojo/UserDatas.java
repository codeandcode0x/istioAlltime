package com.msa.ticket.frontend.pojo;

import com.msa.ticket.frontend.model.User;

import java.util.List;

public class UserDatas {
    private Integer code;
    private List<User> data;

    public Integer getCode() {
        return code;
    }

    public void setCode(Integer code) {
        this.code = code;
    }

    public List<User> getData() {
        return data;
    }

    public void setData(List<User> data) {
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


