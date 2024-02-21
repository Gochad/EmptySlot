interface LoginResponse {
    accessToken: string;
}

class AuthService {
    private static apiUrl: string = 'http://be:8080';

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
        localStorage.setItem('accessToken', data.accessToken);

        return data;
    }

    static logout(): void {
        localStorage.removeItem('accessToken');
    }
}

export default AuthService;