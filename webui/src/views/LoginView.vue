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
                console.log(response.data);
                localStorage.setItem('userID', response.data.user['userID']);
                localStorage.setItem('username', response.data.user['username']);
                localStorage.setItem('propic64', response.data.user['userPropic64']);
                localStorage.setItem('token', response.data.token)

                this.$router.push("/home");
            }).catch((error) => {
                localStorage.setItem('errorMessage', error.response.data);
            }
            );}
        },
    beforeMount() {
    },
    mounted() {
    },
    }

</script>

<template>
    <div class="login-container">
        <div class="top-login-container">
            <span class="top-container-login-title"> Login </span>
        </div>
        <div class="form-container-login">
            <span class="form-text-container-login">Username</span>
            <input type="text" name="username-form" spellcheck="false" v-model="username">
            <span class="form-text-container-login-error">{{ errorMsg }}</span>
        </div>
        <div class="bottom-login-container">
            <button @click="doLogin"> Login </button>
        </div>

    </div>
</template>