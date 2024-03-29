import { NextRequest, NextResponse } from "next/server";

export const mockLoginDetails: Record<string, { authToken: string, userName: string }> = {
  alex: {
    authToken: "123",
    userName: "alex",
  },
  jason: {
    authToken: "234",
    userName: "jason",
  },
  marie: {
    authToken: "345",
    userName: "marie",
  },
};

export const mockCoinDetails: Record<string, { balance: number }> = {
    alex: {
        balance: 100,
    },
    jason: {
        balance: 200,
    },
    marie: {
        balance: 300,
    },
};

export async function GET(_: NextRequest, { params }: { params: { userID: string } }) {
    const authDetails = mockLoginDetails[params.userID];
    return NextResponse.json({ data: mockCoinDetails[authDetails.userName] });
};