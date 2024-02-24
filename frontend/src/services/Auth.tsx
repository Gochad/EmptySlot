interface LoginResponse {
    token: string;
}

class AuthService {
    private static apiUrl: string = 'http://localhost:8080';

    static async login(email: string, password: string): Promise<LoginResponse> {
        const response = await fetch(`${this.apiUrl}/login`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ email, password }),
        });


        if (!response.ok) {
            throw new Error('login error');
        }

        const data: LoginResponse = await response.json();
        localStorage.setItem('token', data.token);

        return data;
    }

    static logout(): void {
        localStorage.removeItem('token');
    }
}

export default AuthService;