import pandas as pd
from sklearn.model_selection import train_test_split

def main():
    # 1. Load the Excel file (adjust path if needed)
    df = pd.read_excel("Folds5x2_pp.xlsx", engine="openpyxl")

    # 2. Split into train (80%) and test (20%)
    train_df, test_df = train_test_split(df, test_size=0.20, random_state=42, shuffle=True)

    # 3. Write out CSVs
    train_df.to_csv("train.csv", index=False)
    test_df.to_csv("test.csv", index=False)

    print(f"Saved train.csv ({len(train_df)} rows) and test.csv ({len(test_df)} rows).")

if __name__ == "__main__":
    main()
