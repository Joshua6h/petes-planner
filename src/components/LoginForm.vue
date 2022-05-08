<template>
    <v-form @submit.prevent="login">
        <v-text-field label="Username" autofocus outlined v-model="username"
                :color="this.$store.state.accent" background-color="white"></v-text-field>
        <v-text-field label="Password" type="password" outlined v-model="password"
                :color="this.$store.state.accent" background-color="white"></v-text-field>
        <router-link to="/reset-password">Forgot password?</router-link>
        <br>
        <v-btn type="submit" class="mt-5 mb-8">Login</v-btn>
        <br>
        <small>No account yet? <router-link to="/sign-up">Sign Up</router-link></small>
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
                this.$emit('raiseError', '');
            }
            else{
                this.$emit('raiseError', 'No fields can be left empty');
                return;
            }
            
            const info = {username: this.username, password: this.password};
            this.$store.dispatch('AUTH_REQUEST', info).then(() => {
                this.$router.push('/');
            }).catch((err) => {
                if (err.response) {
                    // Request made and server responded
                    this.$emit('raiseError', err.response.data);
                }
            })
            .finally(() => {
                this.$store.dispatch('GET_PROFILE');
            })
        }
    }
}
</script>