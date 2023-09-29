if File.exist?("familytree.ged")
    file = File.open("familytree.ged", "r")
    writefile = File.new("familytree.pl", "w+")
    idbase = []
    namebase = []
    namestr = ""
    lines = file.readlines
    lines.each do |x|
        if x[2] == '@' && x[3] == 'I'
            i = 4
            r = ""
            while x[i] != '@'
                r += x[i]
                i += 1
            end
            id = r.to_i
            idbase << id
        end
        if x[2] == 'N' && x.length > 6 && x[0] != '2'
            t = 0
            for i in (7...x.length) do
                if x[i + 1] == '/'
                    t = i
                    break
                end
                namestr << x[i]
            end
            if t != 0
                for j in (i + 3 ...x.length - 1) do
                    namestr << x[j]
                end
            end
            namebase << namestr
            namestr = ""
        end
        if x[2] == '@' && x[3] == 'F'
            break
        end
    end
    namebase.delete_at(0)
    base = {}
    i = 0
    idbase.each do |x|
        base[x] = namebase[i]
        i += 1
    end
    father = ""
    mother = ""
    child = ""
    namebase.clear
    idbase.clear
    lines.each do |x|
        if x[2] == 'H' && x[3] == 'U'
            i = 9
            r = ""
            while x[i] != '@'
                r += x[i]
                i += 1
            end
            id = r.to_i
            f = id
            father = base[f]
        end
        if x[2] == 'W' && x[3] == 'I' && x[4] == 'F' && x[5] == 'E'
            i = 9
            r = ""
            while x[i] != '@'
                r += x[i]
                i += 1
            end
            id = r.to_i
            m = id
            mother = base[m]
        end
        if x[2] == 'C' && x[3] == 'H' && x[4] == 'I' && x[5] == 'L'
            i = 9
            r = ""
            while x[i] != '@'
                r += x[i]
                i += 1
            end
            id = r.to_i
            ch = id
            child = base[ch]
            writefile.puts "parents('#{child}', '#{father}', '#{mother}')."
        end
    end
    lines.clear
    base.clear
    file.close
    writefile.close
else
    puts("File familytree.ged not found")
end
