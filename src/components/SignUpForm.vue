<template>
    <v-form @submit.prevent="signUp">
        <form-text-input label="Username" @contentUpdate="setUsername($event)"></form-text-input>
        <form-text-input label="Password" isPassword @contentUpdate="setPassword($event)"></form-text-input>
        <form-text-input label="Re-enter Password" isPassword @contentUpdate="setSecondPassword($event)"></form-text-input>
        <v-btn type="submit">Sign Up</v-btn>
    </v-form>
</template>

<script>
import FormTextInput from '@/components/FormTextInput.vue';

export default {
    name: 'SignUpForm',
    data() {
        return {
            username: '',
            password: '',
            secondPassword: ''
        }
    },
    components: {
        FormTextInput
    },
    methods: {
        signUp() {
            if(this.username && this.password && this.secondPassword) {
                if(this.password === this.secondPassword){
                    console.log("signed up");
                    console.log(this.username);
                    console.log(this.password);
                    console.log(this.secondPassword);
                    this.$emit('raiseError', '');
                }
                else{
                    this.$emit('raiseError', 'The passwords do not match');
                }
            }
            else{
                this.$emit('raiseError', 'No fields can be left empty');
            }
        },
        setUsername(str) {
            this.username = str;
        },
        setPassword(str) {
            this.password = str;
        },
        setSecondPassword(str) {
            this.secondPassword = str;
        }
    }
}
</script>