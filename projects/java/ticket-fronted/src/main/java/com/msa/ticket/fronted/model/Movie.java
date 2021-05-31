package com.msa.ticket.fronted.model;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

@JsonIgnoreProperties(ignoreUnknown = true)
public class Movie {

    private Integer id;
    private String name;
    private String actors;
    private String image;
    private String mtype;
    private String minfo;
    private String mtime;
    private String CreatedAt;
    private String UpdatedAt;
    private String DeletedAt;
//    private Integer stars;

    public Movie() {
    }

    public Integer getId() {
        return id;
    }

    public void setId(Integer id) {
        this.id = id;
    }

    public String getMtype() {
        return mtype;
    }

    public void setMtype(String mtype) {
        this.mtype = mtype;
    }

    public String getMinfo() {
        return minfo;
    }

    public void setMinfo(String minfo) {
        this.minfo = minfo;
    }

    public String getMtime() {
        return mtime;
    }

    public void setMtime(String mtime) {
        this.mtime = mtime;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getActors() {
        return actors;
    }

    public void setActors(String actors) {
        this.actors = actors;
    }

    public String getImage() {
        return image;
    }

    public void setImage(String image) {
        this.image = image;
    }

    public String getCreatedAt() {
        return CreatedAt;
    }

    public void setCreatedAt(String createdAt) {
        CreatedAt = createdAt;
    }

    public String getUpdatedAt() {
        return UpdatedAt;
    }

    public void setUpdatedAt(String updatedAt) {
        UpdatedAt = updatedAt;
    }

    public String getDeletedAt() {
        return DeletedAt;
    }

    public void setDeletedAt(String deletedAt) {
        DeletedAt = deletedAt;
    }

    //    public Integer getStars() {
//        return stars;
//    }
//
//    public void setStars(Integer stars) {
//        this.stars = stars;
//    }

    @Override
    public String toString() {
        return "Value{" +
                "id=" + id +
                ", name='" + name + '\'' +
                ", actors='" + actors + '\'' +
                ", image='" + image + '\'' +
                ", mtype='" + mtype + '\'' +
                ", minfo='" + minfo + '\'' +
                ", mtime='" + mtime + '\'' +
                ", CreatedAt='" + CreatedAt + '\'' +
                ", UpdatedAt='" + UpdatedAt + '\'' +
                ", DeletedAt='" + DeletedAt + '\'' +
                '}';
    }
}