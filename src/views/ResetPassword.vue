<template>
    <span>
        <h1>Reset Password Here</h1>
        <v-container class="box pa-5">
            <h2 v-if="email != ''">Enter the code you received</h2>
            <v-text-field v-if="email == ''" v-model="input" label="Email" outlined autofocus :color="this.$store.state.accent"></v-text-field>
            <v-text-field v-else-if="email != ''" v-model="input" label="Code" outlined autofocus :color="this.$store.state.accent"></v-text-field>
            <v-btn v-if="email == ''" :disabled="input == '' ? true : false" @click="checkEmail">Continue</v-btn>
            <v-btn v-if="email != ''" @click="resetEmail" class="mr-3">Back</v-btn>
            <v-btn v-if="email != ''" :disabled="input == '' ? true : false" @click="submitCode">Submit</v-btn>
            <br class="mb-10">
            <router-link to="/login">Back to login</router-link>
        </v-container>
    </span>
</template>

<script>
export default {
    name: 'ResetPassword',
    data() {
        return {
            input: '',
            email: ''
        }
    },
    methods: {
        checkEmail() {
            console.log("Yes");
            if(this.input != '') {
                this.email = this.input;
                this.input = '';
            }
        },
        resetEmail() {
            this.email = '';
            this.input = '';
        },
        submitCode() {
            this.input = '';
        }
    }
}
</script>

<style scoped>
.box {
    max-width: 30%;
}
</style>