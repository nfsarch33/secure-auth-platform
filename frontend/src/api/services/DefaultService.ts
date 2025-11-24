/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { AuthResponse } from '../models/AuthResponse';
import type { SignInRequest } from '../models/SignInRequest';
import type { SignUpRequest } from '../models/SignUpRequest';
import type { User } from '../models/User';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class DefaultService {
    /**
     * Register a new user
     * @param requestBody
     * @returns AuthResponse User created successfully
     * @throws ApiError
     */
    public static signUp(
        requestBody: SignUpRequest,
    ): CancelablePromise<AuthResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/auth/signup',
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                400: `Invalid input`,
                409: `User already exists`,
            },
        });
    }
    /**
     * Authenticate a user
     * @param requestBody
     * @returns AuthResponse Authentication successful
     * @throws ApiError
     */
    public static signIn(
        requestBody: SignInRequest,
    ): CancelablePromise<AuthResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/auth/signin',
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                401: `Invalid credentials`,
            },
        });
    }
    /**
     * Sign out a user
     * @returns any Sign out successful
     * @throws ApiError
     */
    public static signOut(): CancelablePromise<any> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/auth/signout',
            errors: {
                401: `Unauthorized`,
            },
        });
    }
    /**
     * Get current user profile
     * @returns User User profile retrieved successfully
     * @throws ApiError
     */
    public static getCurrentUser(): CancelablePromise<User> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/me',
            errors: {
                401: `Unauthorized`,
            },
        });
    }
}
