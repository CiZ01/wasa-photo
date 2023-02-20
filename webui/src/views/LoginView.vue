<script>
export default {
    data() {
        return {
            username: "",
            errorMsg: "",
        }
    },
    methods: {
        doLogin() {
            this.$axios.post('/session', { username: this.username }).then((response) => {
                console.log(response);
                localStorage.userID = response.data.User.userID;
                localStorage.username = response.data.User.username;
                localStorage.token = response.data.Token;

                this.$router.push("/home");
            }).catch((error) => {
                console.log(error);
                this.errorMsg = error.response.data;
            }
            );}
        },
    }

</script>

<template>
    <div class="login-container">
        <div class="top-login-container">
            <span class="top-container-title">Wasa Photo</span>
            <span class="top-container-login-title"> Login </span>
        </div>
        <div class="form-container-login">
            <span class="form-text-container-login">Username</span>
            <input type="text" placeholder="" name="username-form" spellcheck="false" v-model="username">
            <span class="form-text-container-login-error">{{ errorMsg }}</span>
        </div>
        <div class="bottom-login-container">
            <button @click="doLogin"> Login </button>
        </div>

    </div>
</template>