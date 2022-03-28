<template>
    <span>
        <h1>Reset Password Here</h1>
        <h2>Get security question to ask here</h2>
        <v-container class="box pa-5">
            <h2 v-if="username != ''">This is the question</h2>
            <v-text-field v-if="username == ''" v-model="input" label="Username" outlined autofocus :color="this.$store.state.accent"></v-text-field>
            <v-text-field v-else-if="username != ''" v-model="input" label="Answer" outlined autofocus :color="this.$store.state.accent"></v-text-field>
            <v-btn v-if="username == ''" :disabled="input == '' ? true : false" @click="checkUsername">Continue</v-btn>
            <v-btn v-if="username != ''" @click="resetUsername" class="mr-3">Back</v-btn>
            <v-btn v-if="username != ''" :disabled="input == '' ? true : false" @click="submitAnswer">Submit</v-btn>
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
            username: ''
        }
    },
    methods: {
        checkUsername() {
            console.log("Yes");
            if(this.input != '') {
                this.username = this.input;
                this.input = '';
            }
        },
        resetUsername() {
            this.username = '';
            this.input = '';
        },
        submitAnswer() {
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