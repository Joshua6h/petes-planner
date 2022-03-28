<template>
    <v-form @submit.prevent="login">
        <form-text-input label="Username" @contentUpdate="updateUsername($event)"></form-text-input>
        <form-text-input label="Password" isPassword @contentUpdate="updatePassword($event)"></form-text-input>
        <router-link to="/reset-password">Forgot password?</router-link>
        <br>
        <v-btn type="submit" class="mt-5 mb-8">Login</v-btn>
        <br>
        <small>No account yet? <router-link to="/sign-up">Sign Up</router-link></small>
    </v-form>
</template>

<script>
import FormTextInput from '@/components/FormTextInput.vue';

export default {
    name: 'LoginForm',
    components: {
        FormTextInput
    },
    data() {
        return {
            username: '',
            password: '',
            error: ''
        }
    },
    methods: {
        login() {
            console.log('Login button was pressed'); //These are not staying like this, obviously
            if(this.username && this.password){
                console.log(`${this.username} ${this.password}`);
                this.$emit('raiseError', '');
            }
            else{
                this.$emit('raiseError', 'No fields can be left empty');
            }
        },
        updateUsername(str) {
            this.username = str;
        },
        updatePassword(str) {
            this.password = str;
        },
        changeForm() {
            this.$store.dispatch('switchForm', 'Sign In Page')
        }
    }
}
</script>