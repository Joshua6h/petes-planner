<template>
    <span>
        <h1>Password Reset</h1>
        <v-container class="box pa-5">
            <h2 v-if="emailEntered" class='green--text'>An email has been sent to this email address</h2>
            <v-text-field v-if="!emailEntered" v-model="input" label="Email" outlined autofocus :color="this.$store.state.accent"></v-text-field>
            <v-btn v-if="!emailEntered" :disabled="input == '' ? true : false" @click="checkEmail">Continue</v-btn>
            <br class="mb-10">
            <v-btn small to="/login">Back to login</v-btn>
        </v-container>
    </span>
</template>

<script>
export default {
    name: 'ResetForm',
    data() {
        return {
            input: '',
            emailEntered: false,
            newPassword: ''
        }
    },
    methods: {
        checkEmail() {
            if(this.input != '') {
                let data = {email: this.input};
                this.$store.dispatch('SEND_EMAIL', data)
                .then(() => {
                    this.input = '';
                    this.emailEntered = true;
                })
                .catch(() => {
                    this.input = '';
                })
            }
        }
    }
}
</script>

<style scoped>
.box {
    max-width: 30%;
}
</style>