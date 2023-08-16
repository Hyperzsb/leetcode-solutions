class Solution {
    private String normalize(String raw) {
        String[] parts = raw.split("@");
        String local = parts[0], host = parts[1];

        local = local.replaceAll("[.]", "");
        if (local.indexOf("+") != -1) {
            local = local.substring(0, local.indexOf("+"));
        }

        return local + "@" + host;
    }

    public int numUniqueEmails(String[] emails) {
        Set<String> hs = new HashSet<>();
        for (String email : emails) {
            hs.add(normalize(email));
        } 

        System.out.println(hs);

        return hs.size();
    }
}

