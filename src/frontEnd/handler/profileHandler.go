package handler

import (
	"net/http"
    "frontEnd/server"
    //"backEnd/cmd"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request, srv *server.Server){
    //var buttons []FollowButton
    token := r.FormValue("Auth")
    var user User
    user.token = token
    var profile ProfilePage
    if (true) {
        //action := r.FormValue("FollowOrNot")
        //target := r.FormValue("Target")
        deleteAccount := r.FormValue("delete")

        if deleteAccount != "" {
        	//args := cmd.DeleteUserArgs{ Token:token }
            //var reply DeleteReply
            //srv.SrvClient.Call("Server.DeleteUser", args, &reply)
            http.Redirect(w, r, "/login/", http.StatusFound)
            return
        }

        /*if action == "Unfollow" {
            user.UnFollow(srv.users[target])
        }else if(action == "Follow"){
            user.Follow(srv.users[target])
        }

        for u := range(srv.users){
            _, ok := user.following[u]
            if u != user.Username {
                if ok {
                    buttons = append(buttons, FollowButton{Name:u, Action:"Unfollow", User:user})
                } else {
                    buttons = append(buttons, FollowButton{Name:u, Action:"Follow", User:user})
                }
            }
        }
        profile = ProfilePage{User:user, FollowList:buttons}
    } else {
        http.Redirect(w, r, "/login/", http.StatusFound)
        return*/
    }
    server.RenderTemplate(w, srv, "profile", profile)
}

