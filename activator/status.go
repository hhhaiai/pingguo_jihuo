package activator

import (
    "fmt"
)

const (
    ACTIVATE_SUCCESS                = 0
    ACTIVATE_FAILED                 = 1
    ACTIVATE_PROXY_DISCONNECTED     = 2
    ACTIVATE_EMAIL_NOT_FOUND        = 3
    ACTIVATE_MAIL_LOGIN_FAILED      = 4
    ACTIVATE_SESSION_TIMEOUT        = 5
    ACTIVATE_LINK_EXPIRED           = 6
    ACTIVATE_EMAIL_RESENT           = 7
    ACTIVATE_DELAY_PROCESS          = 8
    ACTIVATE_PROXY_INVALID          = 9
    ACTIVATE_MAIL_CONNECT_FAILED    = 10
    ACTIVATE_NETWORK_ERROR          = 11
    ACTIVATE_ACCOUNT_INVALID        = 12
    ACTIVATE_CANT_FINISH_REQUEST    = 13
    ACTIVATE_CANT_ACTIVATE          = 14
    ACTIVATE_PASSWORD_ERROR         = 15
    ACTIVATE_INVALID_PASSWORD_TOKEN = 16
    ACTIVATE_ACCOUNT_BANNED         = 17
)

type ActivateResult int

var activateResultText = map[ActivateResult]string{
    ACTIVATE_SUCCESS                : "ACTIVATE_SUCCESS",
    ACTIVATE_FAILED                 : "ACTIVATE_FAILED",
    ACTIVATE_PROXY_DISCONNECTED     : "ACTIVATE_PROXY_DISCONNECTED",
    ACTIVATE_EMAIL_NOT_FOUND        : "ACTIVATE_EMAIL_NOT_FOUND",
    ACTIVATE_MAIL_LOGIN_FAILED      : "ACTIVATE_MAIL_LOGIN_FAILED",
    ACTIVATE_SESSION_TIMEOUT        : "ACTIVATE_SESSION_TIMEOUT",
    ACTIVATE_LINK_EXPIRED           : "ACTIVATE_LINK_EXPIRED",
    ACTIVATE_EMAIL_RESENT           : "ACTIVATE_EMAIL_RESENT",
    ACTIVATE_DELAY_PROCESS          : "ACTIVATE_DELAY_PROCESS",
    ACTIVATE_PROXY_INVALID          : "ACTIVATE_PROXY_INVALID",
    ACTIVATE_MAIL_CONNECT_FAILED    : "ACTIVATE_MAIL_CONNECT_FAILED",
    ACTIVATE_NETWORK_ERROR          : "ACTIVATE_NETWORK_ERROR",
    ACTIVATE_ACCOUNT_INVALID        : "ACTIVATE_ACCOUNT_INVALID",
    ACTIVATE_CANT_FINISH_REQUEST    : "ACTIVATE_CANT_FINISH_REQUEST",
    ACTIVATE_CANT_ACTIVATE          : "ACTIVATE_CANT_ACTIVATE",
    ACTIVATE_PASSWORD_ERROR         : "ACTIVATE_PASSWORD_ERROR",
    ACTIVATE_INVALID_PASSWORD_TOKEN : "ACTIVATE_INVALID_PASSWORD_TOKEN",
    ACTIVATE_ACCOUNT_BANNED         : "ACTIVATE_ACCOUNT_BANNED",
}

func (self ActivateResult) String() string {
    text, ok := activateResultText[self]
    if ok {
        return text
    }

    return fmt.Sprintf("%d", self)
}
