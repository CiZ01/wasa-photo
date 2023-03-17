<script>
export default {
    data() {
        return {
            username: "",
            errorMsg: "",

            isLoading: false,
        }
    },
    methods: {
        async doLogin() {
            this.isLoading = true;
            try {
                let response = await this.$axios.post('/session', {
                    username: this.username,
                });
                localStorage.userID = response.data.user.userID;
                localStorage.username = response.data.user.username;
                localStorage.propic64 = response.data.user.userPropic64;
                localStorage.token = response.data.token;
                this.$router.push("/home");
            } catch (e) {
                this.errorMsg = e.toString();
                document.getElementsByTagName("input")[0].style.outline = "auto";
                document.getElementsByTagName("input")[0].style.outlineColor = "red";
            };
            this.isLoading = false;
        }
    },
    mounted() {
        localStorage.clear();
    },
}

</script>

<template>
    <LoadingSpinner :loading="isLoading"></LoadingSpinner>
    <ErrorMsg v-if="errorMsg" :msg="errorMsg" @close-error="errorMsg = ''"></ErrorMsg>
    <div class="login-container">
        <div class="top-login-container">
            <span class="top-container-login-title"> Login </span>
        </div>
        <div class="form-container-login">
            <span class="form-text-container-login">Username</span>
            <input :on-submit="doLogin" type="text" name="username-form" spellcheck="false" v-model="username"
                maxlength="13">
        </div>
        <div class="bottom-login-container">
            <button @click="doLogin"> Login </button>
        </div>

    </div>
</template>