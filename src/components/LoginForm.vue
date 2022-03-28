<template>
    <v-form @submit.prevent='login'>
        <v-text-field color='#EE4E34' background-color='white' label='Username' v-model='username' outlined autofocus></v-text-field>
        <v-text-field color='#EE4E34' background-color='white' label='Password' v-model='password' outlined type='password'></v-text-field>
        <router-link to='/reset-password'>Forgot password?</router-link>
        <br>
        <v-btn type='submit' class='mt-5 mb-8'>Login</v-btn>
        <div>
            <small>No account yet? <router-link to='/'>Sign Up</router-link></small>
        </div>
        <v-container class='error my-5' v-if="error != ''">
            <h5>{{error}}</h5>
        </v-container>
    </v-form>
</template>

<script>
export default {
    name: 'LoginForm',
    data() {
        return {
            username: '',
            password: '',
            error: ''
        }
    },
    methods: {
        login: function() {
            console.log('Login button was pressed'); //These are not staying like this, obviously
            if(this.username && this.password){
                console.log(`${this.username} ${this.password}`);
                this.error = '';
            }
            else{
                this.error = 'No fields can be left empty';
            }
        }
    }
}
</script>

<style scoped>
div {
    font-size: 15px;
}
</style>