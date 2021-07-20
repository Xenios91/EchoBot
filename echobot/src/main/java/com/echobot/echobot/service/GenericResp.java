package com.echobot.echobot.service;

@FunctionalInterface
public interface GenericResp<T, U> {
    public abstract T respond(U toUpdate);
}
