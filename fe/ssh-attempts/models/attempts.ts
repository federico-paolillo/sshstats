export interface LoginAttempt {
  username: string
  count: number
}

export interface LoginAttempts {
  attempts: LoginAttempt[]
}