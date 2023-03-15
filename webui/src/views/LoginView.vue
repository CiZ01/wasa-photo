<script>
export default {
    data() {
        return {
            username: "",
            errorMsg: "",
        }
    },
    methods: {
        async doLogin() {
            try {
                let response = await this.$axios.post('/session', {
                    username: this.username,
                });
                console.log(response.data);
                localStorage.userID = response.data.user.userID;
                localStorage.username = response.data.user.username;
                localStorage.propic64 = response.data.user.userPropic64;
                console.log("ciao");
                localStorage.token = response.data.token;
                console.log("ciao2");
                this.$router.push("/home");

                console.log("ciao3");
            } catch (e) {
                console.log(e);
                this.errorMsg = e.toString();
                document.getElementsByTagName("input")[0].style.outline = "auto";
                document.getElementsByTagName("input")[0].style.outlineColor = "red";
            }
        },
    },
    beforeMount(){
        localStorage.clear();
    }

}
</script>

<template>
    <ErrorMsg v-if="errorMsg" :msg="errorMsg" @close-error="errorMsg = ''"></ErrorMsg>
    <div class="login-container">
        <div class="top-login-container">
            <span class="top-container-login-title"> Login </span>
        </div>
        <div class="form-container-login">
            <span class="form-text-container-login">Username</span>
            <input type="text" name="username-form" spellcheck="false" v-model="username">
        </div>
        <div class="bottom-login-container">
            <button @click="doLogin" type="submit"> Login </button>
        </div>
    </div>
</template>